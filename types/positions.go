package types

type PositionEvent struct {
	Id            string
	MarginAccount string
	MarketId      string
	Size          string
	EntryPrice    string
	Leverage      string
	EntryTime     string
	Side          string
	TpOrderId     string
	SlOrderId     string
	Status        string
}

func NewPositionEvent(
	positionId string,
	marginAccount string,
	marketId string,
	size string,
	entryPrice string,
	leverage string,
	entryTime string,
	side string,
	tpOrderId string,
	slOrderId string,
	status string,
) PositionEvent {
	return PositionEvent{
		Id:            positionId,
		MarginAccount: marginAccount,
		MarketId:      marketId,
		Size:          size,
		EntryPrice:    entryPrice,
		Leverage:      leverage,
		EntryTime:     entryTime,
		Side:          side,
		TpOrderId:     tpOrderId,
		SlOrderId:     slOrderId,
		Status:        status,
	}
}
