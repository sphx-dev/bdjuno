package database

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
)

func (db *Db) SaveOrderUpdate(event types.OrderEvent) error {
	query := `
		INSERT INTO order_event (
			id, account_id, side, quantity, price, order_type, trigger_price, leverage, timestamp, market_id, status
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
		ON CONFLICT (id) DO UPDATE SET
			account_id = EXCLUDED.account_id,
			side = EXCLUDED.side,
			quantity = EXCLUDED.quantity,
			price = EXCLUDED.price,
			order_type = EXCLUDED.order_type,
			trigger_price = EXCLUDED.trigger_price,
			leverage = EXCLUDED.leverage,
			timestamp = EXCLUDED.timestamp,
			market_id = EXCLUDED.market_id,
			status = EXCLUDED.status;
	`

	_, err := db.SQL.Exec(
		query,
		event.Id,
		event.AccountId,
		event.Side,
		event.Quantity,
		event.Price,
		event.OrderType,
		event.TriggerPrice,
		event.Leverage,
		event.Timestamp,
		event.MarketId,
		event.Status,
	)

	if err != nil {
		return fmt.Errorf("failed to save order update: %w", err)
	}
	return nil
}
