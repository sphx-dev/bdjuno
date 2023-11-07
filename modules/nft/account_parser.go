package nft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CoreumFoundation/coreum/v3/x/nft"
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
