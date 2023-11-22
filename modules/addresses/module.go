package addresses

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/juno/v5/modules"
	junomessages "github.com/forbole/juno/v5/modules/messages"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/asset/ft module
type Module struct {
	cdc           codec.Codec
	db            *database.Db
	messageParser junomessages.MessageAddressesParser
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser,
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "addresses"
}
