package consensus

import (
	"sync"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/juno/v5/modules"
	tmtypes "github.com/cometbft/cometbft/types"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.BlockModule              = &Module{}
)

// Module implements the consensus utils
type Module struct {
	db *database.Db

	mu                sync.Mutex
	realProposers     map[int64]tmtypes.Address
	expectedProposers map[int64]tmtypes.Address
}

// NewModule builds a new Module instance
func NewModule(db *database.Db) *Module {
	return &Module{
		db:                db,
		realProposers:     map[int64]tmtypes.Address{},
		expectedProposers: map[int64]tmtypes.Address{},
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}
