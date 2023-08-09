package nft

import (
	"github.com/CoreumFoundation/coreum/v2/x/nft"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/assetft module
func MessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) { //nolint:gocritic //bdjuno style
	case *nft.MsgSend:
		return []string{msg.Sender, msg.Receiver}, nil
	}

	return nil, nil
}
