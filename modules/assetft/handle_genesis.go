package assetft

import (
	"encoding/json"
	"fmt"

	assetfttypes "github.com/CoreumFoundation/coreum/v2/x/asset/ft/types"
	"github.com/forbole/bdjuno/v3/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "assetft").Msg("parsing genesis")

	// Read the genesis state
	var genState assetfttypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[assetfttypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling assetft state: %s", err)
	}

	// Save the params
	err = m.db.SaveAssetFTParams(types.NewAssetFTParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis assetft params: %s", err)
	}

	return nil
}
