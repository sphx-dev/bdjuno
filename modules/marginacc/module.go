package positions

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module            = &Module{}
	_ modules.TransactionModule = &Module{}
)

// Module represents the x/feemodel module
type Module struct {
	cdc codec.Codec
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "marginacc"
}
