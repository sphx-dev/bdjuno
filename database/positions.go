package database

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
)

func (db *Db) SavePositionUpdate(event types.PositionEvent) error {
	query := `
		INSERT INTO position_event (
			id, margin_account, market_id, size, entry_price, leverage, entry_time, side, tp_order_id, sl_order_id, status
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
		ON CONFLICT (id) DO UPDATE SET
			margin_account = EXCLUDED.margin_account,
			market_id = EXCLUDED.market_id,
			size = EXCLUDED.size,
			entry_price = EXCLUDED.entry_price,
			leverage = EXCLUDED.leverage,
			entry_time = EXCLUDED.entry_time,
			side = EXCLUDED.side,
			tp_order_id = EXCLUDED.tp_order_id,
			sl_order_id = EXCLUDED.sl_order_id,
			status = EXCLUDED.status;
	`

	_, err := db.SQL.Exec(
		query,
		event.Id,
		event.MarginAccount,
		event.MarketId,
		event.Size,
		event.EntryPrice,
		event.Leverage,
		event.EntryTime,
		event.Side,
		event.TpOrderId,
		event.SlOrderId,
		event.Status,
	)

	if err != nil {
		return fmt.Errorf("failed to save position update: %w", err)
	}
	return nil
}
