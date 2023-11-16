package assetnft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	assetnfttypes "github.com/CoreumFoundation/coreum/v3/x/asset/nft/types"
)

// MessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/assetft module
func MessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *assetnfttypes.MsgIssueClass:
		return []string{msg.Issuer}, nil
	case *assetnfttypes.MsgMint:
		accounts := []string{msg.Sender}
		if msg.Recipient != "" {
			accounts = append(accounts, msg.Recipient)
		}
		return accounts, nil
	case *assetnfttypes.MsgBurn:
		return []string{msg.Sender}, nil
	case *assetnfttypes.MsgAddToWhitelist:
		return []string{msg.Sender, msg.Account}, nil
	case *assetnfttypes.MsgRemoveFromWhitelist:
		return []string{msg.Sender, msg.Account}, nil
	}

	return nil, nil
}
