package wasm

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

// Module represents the x/staking module
type Module struct {
	cdc           codec.Codec
	db            *database.Db
	messageParser junomessages.MessageAddressesParser
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser,
	cdc codec.Codec,
	db *database.Db,
) *Module {
	return &Module{
		messageParser: messageParser,
		cdc:           cdc,
		db:            db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "wasm"
}
