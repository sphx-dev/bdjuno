package addresses

import (
	"encoding/json"
	"strings"

	tmtypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	juno "github.com/forbole/juno/v5/types"
	"github.com/samber/lo"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	addresses, err := m.collectAddresses(msg, tx)
	if err != nil {
		return err
	}

	// marshal the value properly
	bz, err := m.cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	return m.db.SaveMessage(juno.NewMessage(
		tx.TxHash,
		index,
		proto.MessageName(msg),
		string(bz),
		addresses,
		tx.Height,
	))
}

func (m *Module) collectAddresses(msg sdk.Msg, tx *juno.Tx) ([]string, error) {
	// get the involved addresses with general parser first
	messageAddresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		return nil, err
	}

	addresses := make(map[string]struct{})
	for _, address := range messageAddresses {
		addresses[address] = struct{}{}
	}
	// add address from event values
	addBech32EventValues(addresses, tx.Events)
	if err := addBech32MsgValues(addresses, m.cdc, msg); err != nil {
		return nil, err
	}

	return lo.Keys(addresses), nil
}

func addBech32MsgValues(addressSet map[string]struct{}, cdc codec.Codec, msg sdk.Msg) error {
	msgJSON, err := cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	objects := []interface{}{map[string]interface{}{}}
	if err := json.Unmarshal(msgJSON, &objects[0]); err != nil {
		return err
	}

	for len(objects) > 0 {
		object := objects[len(objects)-1]
		objects = objects[:len(objects)-1]

		switch v := object.(type) {
		case map[string]interface{}:
			for _, o := range v {
				objects = append(objects, o)
			}
		case []interface{}:
			objects = append(objects, v...)
		case string:
			if isBech32Address(v) {
				addressSet[v] = struct{}{}
			}
		}
	}

	return nil
}

func addBech32EventValues(addressSet map[string]struct{}, events []tmtypes.Event) {
	for _, ev := range sdk.StringifyEvents(events) {
		for _, attrItem := range ev.Attributes {
			address := strings.Trim(strings.TrimSpace(attrItem.Value), `"`)
			if !isBech32Address(address) {
				continue
			}
			addressSet[address] = struct{}{}
		}
	}
}

func isBech32Address(address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	return err == nil
}
