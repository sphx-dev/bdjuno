package types

type MarketEvent struct {
	MarketId         string
	MarketTicker     string
	MarketBaseAsset  string
	MarketQuoteAsset string
	MarketStatus     string
}

func NewMarketEvent(
	marketID string,
	marketTicker string,
	marketBaseAsset string,
	marketQuoteAsset string,
	marketStatus string,
) MarketEvent {
	return MarketEvent{
		MarketId:         marketID,
		MarketTicker:     marketTicker,
		MarketBaseAsset:  marketBaseAsset,
		MarketQuoteAsset: marketQuoteAsset,
		MarketStatus:     marketStatus,
	}
}
