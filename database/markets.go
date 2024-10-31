package database

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
)

func (db *Db) SaveRegisterMarketEvent(event types.MarketEvent) error {
	query := `
		INSERT INTO market_events (
			market_id, market_ticker, market_base_asset, market_quote_asset, market_status
		) VALUES (
			$1, $2, $3, $4, $5
		)
		ON CONFLICT (market_id) DO UPDATE SET
			market_ticker = EXCLUDED.market_ticker,
			market_base_asset = EXCLUDED.market_base_asset,
			market_quote_asset = EXCLUDED.market_quote_asset,
			market_status = EXCLUDED.market_status;
	`

	_, err := db.SQL.Exec(
		query,
		event.MarketId,
		event.MarketTicker,
		event.MarketBaseAsset,
		event.MarketQuoteAsset,
		event.MarketStatus,
	)

	if err != nil {
		return fmt.Errorf("failed to save market event: %w", err)
	}
	return nil
}
