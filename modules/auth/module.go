package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	authsource "github.com/forbole/bdjuno/v4/modules/auth/source"

	"github.com/forbole/bdjuno/v4/database"

	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/modules/messages"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	cdc            codec.Codec
	db             *database.Db
	source         authsource.Source
	messagesParser messages.MessageAddressesParser
}

// NewModule builds a new Module instance
func NewModule(source authsource.Source, messagesParser messages.MessageAddressesParser, cdc codec.Codec,
	db *database.Db) *Module {
	return &Module{
		messagesParser: messagesParser,
		cdc:            cdc,
		db:             db,
		source:         source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}
