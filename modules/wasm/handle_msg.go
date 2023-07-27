package wasm

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"
	"github.com/samber/lo"
	tmtypes "github.com/tendermint/tendermint/abci/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) { //nolint:gocritic //bdjuno style
	case *wasmtypes.MsgExecuteContract:
		return m.handleMsgExecuteContract(index, cosmosMsg, tx)
	}

	return nil
}

func (m *Module) handleMsgExecuteContract(index int, msg *wasmtypes.MsgExecuteContract, tx *juno.Tx) error {
	// get the involved addresses with general parser first
	addresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		return err
	}

	receivers := findStringEventAttributes(tx.Events, banktypes.EventTypeCoinReceived, banktypes.AttributeKeyReceiver)
	if len(receivers) == 0 {
		return nil
	}

	// we join and then do the Uniq since the receivers might be duplicated
	addresses = lo.Uniq(append(addresses, receivers...))

	// Marshal the value properly
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

func findStringEventAttributes(events []tmtypes.Event, etype, attribute string) []string {
	values := make([]string, 0)
	for _, ev := range sdk.StringifyEvents(events) {
		if ev.Type == etype {
			values = append(values, findAttributes(ev, attribute)...)
		}
	}

	return values
}

func findAttributes(ev sdk.StringEvent, attr string) []string {
	values := make([]string, 0)
	for _, attrItem := range ev.Attributes {
		if attrItem.Key == attr {
			values = append(values, attrItem.Value)
		}
	}

	return values
}
