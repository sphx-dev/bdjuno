package wasm

import (
	"strings"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgInstantiateContract,
		*wasmtypes.MsgInstantiateContract2,
		*wasmtypes.MsgMigrateContract,
		*wasmtypes.MsgExecuteContract:
		return m.handleWasmRelatedAddress(index, cosmosMsg, tx)
	}

	return nil
}

func (m *Module) handleWasmRelatedAddress(index int, msg sdk.Msg, tx *juno.Tx) error {
	// get the involved addresses with general parser first
	messageAddresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		return err
	}

	addresses := make(map[string]struct{})
	for _, address := range messageAddresses {
		addresses[address] = struct{}{}
	}
	// add address from event values
	m.addBech32EventValues(addresses, tx.Events)

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
		lo.Keys(addresses),
		tx.Height,
	))
}

func (m *Module) addBech32EventValues(addressSet map[string]struct{}, events []tmtypes.Event) {
	for _, ev := range sdk.StringifyEvents(events) {
		for _, attrItem := range ev.Attributes {
			address := strings.Trim(strings.TrimSpace(attrItem.Value), `"`)
			if !m.isBech32Address(address) {
				continue
			}
			addressSet[address] = struct{}{}
		}
	}
}

func (m *Module) isBech32Address(address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	return err == nil
}
