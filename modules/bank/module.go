package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/bank/source"

	junomessages "github.com/forbole/juno/v5/modules/messages"

	"github.com/forbole/juno/v5/modules"

	"github.com/CoreumFoundation/coreum/v3/pkg/config/constant"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represents the x/bank module
type Module struct {
	cdc codec.Codec
	db  *database.Db

	messageParser junomessages.MessageAddressesParser
	keeper        source.Source
	baseDenom     string
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser,
	keeper source.Source,
	cdc codec.Codec,
	db *database.Db,
	addressPrefix string,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
		keeper:        keeper,
		baseDenom:     getBaseTokenFromAddressPrefix(addressPrefix),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}

func getBaseTokenFromAddressPrefix(addressPrefix string) string {
	switch addressPrefix {
	case constant.AddressPrefixMain:
		return constant.DenomMain
	case constant.AddressPrefixTest:
		return constant.DenomTest
	case constant.AddressPrefixDev:
		return constant.DenomDev
	default:
		panic("unknown address prefix: " + addressPrefix)
	}
}
