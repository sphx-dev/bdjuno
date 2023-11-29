package bank

import (
	"fmt"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if err := m.UpdateAccountsBalances(index, tx); err != nil {
		return err
	}

	return nil
}

func (m *Module) UpdateAccountsBalances(index int, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	if err := m.updateBalanceForEventType(index, tx, banktypes.EventTypeCoinReceived); err != nil {
		return err
	}

	return m.updateBalanceForEventType(index, tx, banktypes.EventTypeCoinSpent)
}

func (m *Module) updateBalanceForEventType(index int, tx *juno.Tx, eventType string) error {
	accountAttribute := banktypes.AttributeKeySpender
	if eventType == banktypes.EventTypeCoinReceived {
		accountAttribute = banktypes.AttributeKeyReceiver
	}

	block, err := m.db.GetLastBlockHeightAndTimestamp()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	events := FindAllEventsByType(index, tx, eventType)
	type addressDenom struct {
		address string
		denom   string
	}
	addressDenomMap := make(map[addressDenom]interface{})
	for _, event := range events {
		account, err := tx.FindAttributeByKey(event, accountAttribute)
		if err != nil {
			return err
		}

		coinString, err := tx.FindAttributeByKey(event, sdk.AttributeKeyAmount)
		if err != nil {
			return err
		}

		coin, err := sdk.ParseCoinNormalized(coinString)
		if err != nil {
			return err
		}

		// since the main governance token exists in every transaction, we have decided to skip processing
		// that token.
		if coin.Denom == m.baseDenom {
			continue
		}

		addressDenomMap[addressDenom{address: account, denom: coin.Denom}] = true
	}

	for ad := range addressDenomMap {
		storedBalance, found, err := m.db.GetAccountDenomBalance(ad.address, ad.denom)
		if err != nil {
			return err
		}
		if found && storedBalance.Height >= block.Height {
			continue
		}

		quriedBalance, err := m.keeper.GetAccountDenomBalance(ad.address, ad.denom, block.Height)
		if err != nil {
			return err
		}

		if quriedBalance == nil {
			return fmt.Errorf("query balance return nil, account: %s, denom:%s", ad.address, ad.denom)
		}

		if err := m.db.SaveAccountDenomBalance(ad.address, *quriedBalance, block.Height); err != nil {
			return err
		}
	}

	return nil
}

func FindAllEventsByType(index int, tx *juno.Tx, eventType string) []sdk.StringEvent {
	var list []sdk.StringEvent
	for _, ev := range tx.Logs[index].Events {
		if ev.Type == eventType {
			list = append(list, ev)
		}
	}
	return list
}
