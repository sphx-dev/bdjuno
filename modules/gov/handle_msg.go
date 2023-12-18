package gov

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
)

// HandleMsgExec implements modules.AuthzMessageModule
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	// legacy gov
	case *govtypesv1beta1.MsgSubmitProposal:
		msgSubmitProposal := msg.(*govtypesv1beta1.MsgSubmitProposal)
		content, ok := msgSubmitProposal.Content.GetCachedValue().(govtypesv1beta1.Content)
		if !ok {
			return fmt.Errorf("failed to cast govtypesv1beta1.MsgSubmitProposal Content to govtypesv1beta1.Content, message:%s", msg.String())
		}

		return m.handleMsgSubmitProposal(tx, index, &govtypesv1.MsgSubmitProposal{
			Messages:       []*sdktypes.Any{msgSubmitProposal.Content},
			InitialDeposit: msgSubmitProposal.InitialDeposit,
			Proposer:       msgSubmitProposal.Proposer,
			Title:          content.GetTitle(),
			Summary:        content.GetDescription(),
			// v1 attribute
			Metadata: "",
		})

	case *govtypesv1beta1.MsgDeposit:
		msgDeposit := msg.(*govtypesv1beta1.MsgDeposit)
		return m.handleMsgDeposit(tx, &govtypesv1.MsgDeposit{
			ProposalId: msgDeposit.ProposalId,
			Depositor:  msgDeposit.Depositor,
			Amount:     msgDeposit.Amount,
		})

	case *govtypesv1beta1.MsgVote:
		msgVote := msg.(*govtypesv1beta1.MsgVote)

		return m.handleMsgVote(tx, &govtypesv1.MsgVote{
			ProposalId: msgVote.ProposalId,
			Voter:      msgVote.Voter,
			Option:     govtypesv1.VoteOption(msgVote.Option),
			// v1 attribute
			Metadata: "",
		})

	case *govtypesv1beta1.MsgVoteWeighted:
		msgVoteWeighted := msg.(*govtypesv1beta1.MsgVoteWeighted)
		return m.handleMsgVoteWeighted(tx, &govtypesv1.MsgVoteWeighted{
			ProposalId: msgVoteWeighted.ProposalId,
			Voter:      msgVoteWeighted.Voter,
			Options: lo.Map(msgVoteWeighted.Options, func(o govtypesv1beta1.WeightedVoteOption, _ int) *govtypesv1.WeightedVoteOption {
				return &govtypesv1.WeightedVoteOption{
					Option: govtypesv1.VoteOption(o.Option),
					Weight: o.Weight.String(),
				}
			}),
			// v1 attribute
			Metadata: "",
		})

	// v1 gov
	case *govtypesv1.MsgSubmitProposal:
		return m.handleMsgSubmitProposal(tx, index, cosmosMsg)

	case *govtypesv1.MsgDeposit:
		return m.handleMsgDeposit(tx, cosmosMsg)

	case *govtypesv1.MsgVote:
		return m.handleMsgVote(tx, cosmosMsg)

	case *govtypesv1.MsgVoteWeighted:
		return m.handleMsgVoteWeighted(tx, cosmosMsg)
	}

	return nil
}

// handleMsgSubmitProposal allows to properly handle a MsgSubmitProposal
func (m *Module) handleMsgSubmitProposal(tx *juno.Tx, index int, msg *govtypesv1.MsgSubmitProposal) error {
	// Get the proposal id
	event, err := tx.FindEventByType(index, gov.EventTypeSubmitProposal)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeSubmitProposal: %s", err)
	}

	id, err := tx.FindAttributeByKey(event, gov.AttributeKeyProposalID)
	if err != nil {
		return fmt.Errorf("error while searching for AttributeKeyProposalID: %s", err)
	}

	proposalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing proposal id: %s", err)
	}

	// Get the proposal
	proposal, err := m.source.Proposal(tx.Height, proposalID)
	if err != nil {
		if strings.Contains(err.Error(), codes.NotFound.String()) {
			// query the proposal details using the latest height stored in db
			// to fix the rpc error returning code = NotFound desc = proposal x doesn't exist
			block, err := m.db.GetLastBlockHeightAndTimestamp()
			if err != nil {
				return fmt.Errorf("error while getting latest block height: %s", err)
			}
			proposal, err = m.source.Proposal(block.Height, proposalID)
			if err != nil {
				return fmt.Errorf("error while getting proposal: %s", err)
			}
		} else {
			return fmt.Errorf("error while getting proposal: %s", err)
		}
	}

	var addresses []types.Account
	for _, msg := range proposal.Messages {
		var sdkMsg sdk.Msg
		err := m.cdc.UnpackAny(msg, &sdkMsg)
		if err != nil {
			return fmt.Errorf("error while unpacking proposal message: %s", err)
		}

		switch msg := sdkMsg.(type) {
		case *distrtypes.MsgCommunityPoolSpend:
			addresses = append(addresses, types.NewAccount(msg.Recipient))
		case *govtypesv1.MsgExecLegacyContent:
			content, ok := msg.Content.GetCachedValue().(*distrtypes.CommunityPoolSpendProposal)
			if ok {
				addresses = append(addresses, types.NewAccount(content.Recipient))
			}
		}
	}

	err = m.db.SaveAccounts(addresses)
	if err != nil {
		return fmt.Errorf("error while storing proposal recipient: %s", err)
	}

	// Store the proposal
	proposalObj := types.NewProposal(
		proposal.Id,
		proposal.Title,
		proposal.Summary,
		proposal.Metadata,
		msg.Messages,
		proposal.Status.String(),
		*proposal.SubmitTime,
		*proposal.DepositEndTime,
		proposal.VotingStartTime,
		proposal.VotingEndTime,
		msg.Proposer,
	)

	err = m.db.SaveProposals([]types.Proposal{proposalObj})
	if err != nil {
		return err
	}

	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	// Store the deposit
	deposit := types.NewDeposit(proposal.Id, msg.Proposer, msg.InitialDeposit, txTimestamp, tx.TxHash, tx.Height)
	return m.db.SaveDeposits([]types.Deposit{deposit})
}

// handleMsgDeposit allows to properly handle a MsgDeposit
func (m *Module) handleMsgDeposit(tx *juno.Tx, msg *govtypesv1.MsgDeposit) error {
	deposit, err := m.source.ProposalDeposit(tx.Height, msg.ProposalId, msg.Depositor)
	if err != nil {
		return fmt.Errorf("error while getting proposal deposit: %s", err)
	}
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveDeposits([]types.Deposit{
		types.NewDeposit(msg.ProposalId, msg.Depositor, deposit.Amount, txTimestamp, tx.TxHash, tx.Height),
	})
}

// handleMsgVote allows to properly handle a MsgVote
func (m *Module) handleMsgVote(tx *juno.Tx, msg *govtypesv1.MsgVote) error {
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	vote := types.NewVote(msg.ProposalId, msg.Voter, msg.Option, "1.0", txTimestamp, tx.Height)

	err = m.db.SaveVote(vote)
	if err != nil {
		return fmt.Errorf("error while saving vote: %s", err)
	}

	// update tally result for given proposal
	return m.UpdateProposalTallyResult(msg.ProposalId, tx.Height)
}

// handleMsgVoteWeighted allows to properly handle a MsgVoteWeighted
func (m *Module) handleMsgVoteWeighted(tx *juno.Tx, msg *govtypesv1.MsgVoteWeighted) error {
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	for _, option := range msg.Options {
		vote := types.NewVote(msg.ProposalId, msg.Voter, option.Option, option.Weight, txTimestamp, tx.Height)
		err = m.db.SaveVote(vote)
		if err != nil {
			return fmt.Errorf("error while saving weighted vote for address %s: %s", msg.Voter, err)
		}
	}

	// update tally result for given proposal
	return m.UpdateProposalTallyResult(msg.ProposalId, tx.Height)
}
