package types

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
