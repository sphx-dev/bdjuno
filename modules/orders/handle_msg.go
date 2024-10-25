package orders

import (
	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
)

//package types

const (
	OrderModuleEventType = "order"

	EventTypeMsgPlaceOrder = "place_order"

	EventTypeMsgCancelOrder = "cancel_order"

	EventTypeMsgRegisterMarket = "register_market"

	EventTypeModifyPosition = "modify_position"

	EventTypeNewPosition = "new_position"
)

const (
	AttributeKeyOrderID              = "order_id"
	AttributeKeyAccountID            = "account_id"
	AttributeKeyPositionID           = "position_id"
	AttributeKeyOrderType            = "order_type"
	AttributeKeyOrderSide            = "order_side"
	AttributeKeyBaseSize             = "base_size"
	AttributeKeyMarketID             = "market_id"
	AttributeKeyPrice                = "price"
	AttributeKeyQuantity             = "quantity"
	AttributeKeyTriggerPrice         = "trigger_price"
	AttributeKeyTimestamp            = "timestamp"
	AttributeKeyLeverage             = "leverage"
	AttributeKeyIsPostOnly           = "is_post_only"
	AttributeKeyIsReduceOnly         = "is_reduce_only"
	AttributeKeyTimeInForce          = "time_in_force"
	AttributeKeyExpireTime           = "expire_time"
	AttributeKeyMarginAccountAddress = "margin_account_address"
	AttributeKeyQuoteAsset           = "quote_asset"
	AttributeKeyQuoteAmount          = "quote_amount"
	AttributeKeySpendableBalance     = "spendable_balance"

	AttributeKeyMarketTicker     = "ticker"
	AttributeKeyMarketBaseAsset  = "base_asset"
	AttributeKeyMarketQuoteAsset = "quote_asset"
	AttributeKeyMarketStatus     = "status"

	AttributeKeyPositionSize   = "position_size"
	AttributeKeyEntryPrice     = "entry_price"
	AttributeKeyEntryTime      = "entry_time"
	AttributeKeyTpOrderID      = "tp_order_id"
	AttributeKeySlOrderID      = "sl_order_id"
	AttributeKeyPositionStatus = "position_status"
	AttributeKeyPositionSide   = "position_side"

	AttributeKeyError = "error"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	// Iterate through the logs in the transaction
	for _, log := range tx.Logs {
		for _, event := range log.Events {
			// Check if the event type matches the order events we're interested in
			if event.Type == EventTypeMsgPlaceOrder || event.Type == EventTypeMsgCancelOrder {
				// Initialize a map to store the attributes of the event
				attributes := make(map[string]string)
				for _, attr := range event.Attributes {
					attributes[attr.Key] = attr.Value
				}

				// Extract the attributes we're interested in
				orderID := attributes[AttributeKeyOrderID]
				accountID := attributes[AttributeKeyAccountID]
				orderSide := attributes[AttributeKeyOrderSide]
				quantity := attributes[AttributeKeyQuantity]
				price := attributes[AttributeKeyPrice]
				triggerPrice := attributes[AttributeKeyTriggerPrice]
				orderType := attributes[AttributeKeyOrderType]
				timestamp := attributes[AttributeKeyTimestamp]
				leverage := attributes[AttributeKeyLeverage]
				marketID := attributes[AttributeKeyMarketID]

				// Do something with the extracted data (e.g., log it, process it, or store it externally)
				// Example: log the order details
				//fmt.Printf("Order Event: ID=%s, Account=%s, Side=%s, Quantity=%s, Price=%s, MarketID=%s\n",
				//	orderID, accountID, orderSide, quantity, price, marketID, triggerPrice, orderType, timestamp, leverage,
				// )

				m.db.SaveOrderUpdate(types.NewOrderEvent(
					orderID,
					accountID,
					orderSide,
					quantity,
					price,
					marketID,
					triggerPrice,
					orderType,
					timestamp,
					leverage),
				)
			}
		}
	}

	return nil
}
