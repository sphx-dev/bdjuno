package types

type OrderEvent struct {
	Id           string
	AccountId    string
	Side         string
	Quantity     string
	Price        string
	OrderType    string
	TriggerPrice string
	Leverage     string
	Timestamp    string
	MarketId     string
	Status       string
	//Fills []Fill
}

func NewOrderEvent(
	orderID string,
	accountID string,
	orderSide string,
	quantity string,
	price string,
	marketID string,
	triggerPrice string,
	orderType string,
	timestamp string,
	leverage string,
) OrderEvent {
	return OrderEvent{
		Id:           orderID,
		AccountId:    accountID,
		Side:         orderSide,
		Quantity:     quantity,
		Price:        price,
		MarketId:     marketID,
		TriggerPrice: triggerPrice,
		OrderType:    orderType,
		Timestamp:    timestamp,
		Leverage:     leverage,
	}
}
