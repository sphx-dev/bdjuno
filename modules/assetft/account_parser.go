package assetft

import (
	assetfttypes "github.com/CoreumFoundation/coreum/v2/x/asset/ft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/assetft module
func MessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *assetfttypes.MsgIssue:
		return []string{msg.Issuer}, nil
	case *assetfttypes.MsgMint:
		return []string{msg.Sender}, nil
	case *assetfttypes.MsgBurn:
		return []string{msg.Sender}, nil
	case *assetfttypes.MsgFreeze:
		return []string{msg.Sender, msg.Account}, nil
	case *assetfttypes.MsgUnfreeze:
		return []string{msg.Sender, msg.Account}, nil
	case *assetfttypes.MsgGloballyFreeze:
		return []string{msg.Sender}, nil
	case *assetfttypes.MsgGloballyUnfreeze:
		return []string{msg.Sender}, nil
	case *assetfttypes.MsgSetWhitelistedLimit:
		return []string{msg.Sender, msg.Account}, nil
	case *assetfttypes.MsgUpgradeTokenV1:
		return []string{msg.Sender}, nil
	}

	return nil, nil
}
