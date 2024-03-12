package consensus

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v4/modules/actions/logging"
	"github.com/forbole/juno/v5/types"

	"github.com/rs/zerolog/log"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
)

// HandleBlock implements modules.Module
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ []*types.Tx, vals *tmctypes.ResultValidators,
) error {
	err := m.updateBlockTimeFromGenesis(b)
	if err != nil {
		log.Error().Str("module", "consensus").Int64("height", b.Block.Height).
			Err(err).Msg("error while updating block time from genesis")
	}

	logging.BlockRoundSummary.WithLabelValues(strconv.Itoa(int(b.Block.LastCommit.Round))).Observe(1.0)

	m.countProposalsByValidator(b, vals)
	m.measureVotingTimes(b)

	return nil
}

// updateBlockTimeFromGenesis insert average block time from genesis
func (m *Module) updateBlockTimeFromGenesis(block *tmctypes.ResultBlock) error {
	log.Trace().Str("module", "consensus").Int64("height", block.Block.Height).
		Msg("updating block time from genesis")

	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}
	if genesis == nil {
		return fmt.Errorf("genesis table is empty")
	}

	// Skip if the genesis does not exist
	if genesis == nil {
		return nil
	}

	newBlockTime := block.Block.Time.Sub(genesis.Time).Seconds() / float64(block.Block.Height-genesis.InitialHeight)
	return m.db.SaveAverageBlockTimeGenesis(newBlockTime, block.Block.Height)
}

func (m *Module) countProposalsByValidator(block *tmctypes.ResultBlock, vals *tmctypes.ResultValidators) {
	if block.Block.Height != vals.BlockHeight {
		log.Warn().Int64("blockHeight", block.Block.Height).Int64("valHeight", vals.BlockHeight).
			Msg("block height and validator height don't match")
		logging.ValidatorBlockMismatchCounter.Inc()
		return
	}

	// Based on https://github.com/cometbft/cometbft/blob/613d602e48adc723e41bd4c9a9ba37a95a510c7c/types/validator.go#L65
	expectedNextProposer := vals.Validators[0]
	if len(vals.Validators) > 1 {
		for _, v := range vals.Validators[1:] {
			if v.ProposerPriority > expectedNextProposer.ProposerPriority ||
				(v.ProposerPriority == expectedNextProposer.ProposerPriority &&
					bytes.Compare(v.Address, expectedNextProposer.Address) < 0) {
				expectedNextProposer = v
			}
		}
	}

	height := block.Block.Height
	currentProposer := block.Block.ProposerAddress
	nextProposer := expectedNextProposer.Address

	m.mu.Lock()
	defer m.mu.Unlock()

	// This handles the case when block is received in order.
	if expectedProposer, exists := m.expectedProposers[height]; exists {
		delete(m.expectedProposers, height)
		updateProposerMetric(expectedProposer, currentProposer)
	} else {
		m.realProposers[height] = currentProposer
	}

	// This handles the case when block is received out of order.
	if realProposer, exists := m.realProposers[height+1]; exists {
		delete(m.realProposers, height+1)
		updateProposerMetric(nextProposer, realProposer)
	} else {
		m.expectedProposers[height+1] = nextProposer
	}

	// Protection against memory leaks when blocks are missed by the indexer.
	// This is naive approach which might lead to loosing some measures, but it won't break statistics.
	const maxCacheSize = 100
	if len(m.realProposers) > maxCacheSize {
		m.realProposers = map[int64]tmtypes.Address{}
	}
	if len(m.expectedProposers) > maxCacheSize {
		m.expectedProposers = map[int64]tmtypes.Address{}
	}
}

func updateProposerMetric(expected, real tmtypes.Address) {
	var value float64
	if bytes.Equal(expected, real) {
		value = 1.0
	}
	logging.ProposalSummary.WithLabelValues(sdk.ConsAddress(expected).String()).Observe(value)
}

func (m *Module) measureVotingTimes(block *tmctypes.ResultBlock) {
	proposerAddr := block.Block.ProposerAddress.Bytes()

	var proposerVoteTime time.Time
	for _, s := range block.Block.LastCommit.Signatures {
		if bytes.Equal(s.ValidatorAddress.Bytes(), proposerAddr) {
			proposerVoteTime = s.Timestamp
			break
		}
	}
	proposer := sdk.ConsAddress(block.Block.ProposerAddress).String()
	for _, s := range block.Block.LastCommit.Signatures {
		logging.VoteTimeHistogram.WithLabelValues(proposer, sdk.ConsAddress(s.ValidatorAddress).String()).
			Observe(float64(s.Timestamp.Sub(proposerVoteTime)) / float64(time.Millisecond))
	}
}
