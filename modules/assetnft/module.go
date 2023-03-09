package assetnft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v3/database"
	assetnftsource "github.com/forbole/bdjuno/v3/modules/assetnft/source"
	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/asset/nft module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source assetnftsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source assetnftsource.Source,
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
	return "assetnft"
}
