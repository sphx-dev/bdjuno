package assetnft

import (
	"encoding/json"
	"fmt"

	assetnfttypes "github.com/CoreumFoundation/coreum/x/asset/nft/types"
	"github.com/forbole/bdjuno/v3/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "assetnft").Msg("parsing genesis")

	// Read the genesis state
	var genState assetnfttypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[assetnfttypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling assetnft state: %s", err)
	}

	// Save the params
	err = m.db.SaveAssetNFTParams(types.NewAssetNFTParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis assetnft params: %s", err)
	}

	return nil
}
