package auth

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"
	authttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/forbole/bdjuno/v4/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "auth").Msg("parsing genesis")

	accounts, err := GetGenesisAccounts(appState, m.cdc)
	if err != nil {
		return fmt.Errorf("error while getting genesis accounts: %s", err)
	}
	err = m.db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing genesis accounts: %s", err)
	}

	vestingAccounts, err := GetGenesisVestingAccounts(appState, m.cdc)
	if err != nil {
		return fmt.Errorf("error while getting genesis vesting accounts: %s", err)
	}
	err = m.db.SaveVestingAccounts(vestingAccounts)
	if err != nil {
		return fmt.Errorf("error while storing genesis vesting accounts: %s", err)
	}

	var genState authttypes.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[authttypes.ModuleName], &genState); err != nil {
		return err
	}

	err = m.db.SaveAuthParams(types.NewAuthParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis auth params: %s", err)
	}

	return nil
}
