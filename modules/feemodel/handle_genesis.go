package feemodel

import (
	"encoding/json"
	"fmt"

	feemodeltypes "github.com/CoreumFoundation/coreum/x/feemodel/types"
	"github.com/forbole/bdjuno/v3/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "feemodel").Msg("parsing genesis")

	// Read the genesis state
	var genState feemodeltypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[feemodeltypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling feemodel state: %s", err)
	}

	// Save the params
	err = m.db.SaveFeeModelParams(types.NewFeeModelParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis feemodel params: %s", err)
	}

	return nil
}
