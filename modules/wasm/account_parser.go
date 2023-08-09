package wasm

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/assetft module
func MessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *wasmtypes.MsgStoreCode:
		return []string{msg.Sender}, nil
	case *wasmtypes.MsgInstantiateContract:
		return []string{msg.Sender, msg.Admin}, nil
	case *wasmtypes.MsgInstantiateContract2:
		return []string{msg.Sender, msg.Admin}, nil
	case *wasmtypes.MsgExecuteContract:
		return []string{msg.Sender, msg.Contract}, nil
	case *wasmtypes.MsgMigrateContract:
		return []string{msg.Sender, msg.Contract}, nil
	case *wasmtypes.MsgUpdateAdmin:
		return []string{msg.Sender, msg.NewAdmin, msg.Contract}, nil
	case *wasmtypes.MsgClearAdmin:
		return []string{msg.Sender, msg.Contract}, nil
	}

	return nil, nil
}
