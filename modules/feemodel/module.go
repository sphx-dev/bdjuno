package feemodel

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v3/database"
	feemodelsource "github.com/forbole/bdjuno/v3/modules/feemodel/source"
	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/feemodel module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source feemodelsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source feemodelsource.Source,
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "feemodel"
}
