package orders

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	log.Info().Str("module", "orders").
		Int64("height", tx.Height).
		Str("tx", tx.TxHash).
		Msg("HandleTx")
	// Iterate through the events in the transaction directly
	for _, event := range tx.Events {
		// Check if the event type matches the order events we're interested in
		if event.Type == types.EventTypeMsgPlaceOrder || event.Type == types.EventTypeMsgCancelOrder {
			// Initialize a map to store the attributes of the event
			attributes := make(map[string]string)
			for _, attr := range event.Attributes {
				attributes[attr.Key] = attr.Value
			}

			// Extract the attributes we're interested in
			orderID := attributes[types.AttributeKeyOrderID]
			accountID := attributes[types.AttributeKeyAccountID]
			orderSide := attributes[types.AttributeKeyOrderSide]
			quantity := attributes[types.AttributeKeyQuantity]
			price := attributes[types.AttributeKeyPrice]
			triggerPrice := attributes[types.AttributeKeyTriggerPrice]
			orderType := attributes[types.AttributeKeyOrderType]
			timestamp := attributes[types.AttributeKeyTimestamp]
			leverage := attributes[types.AttributeKeyLeverage]
			marketID := attributes[types.AttributeKeyMarketID]

			// Save the order event using the extracted attributes
			err := m.db.SaveOrderUpdate(types.NewOrderEvent(
				orderID,
				accountID,
				orderSide,
				quantity,
				price,
				orderType,
				triggerPrice,
				leverage,
				timestamp,
				marketID,
			))
			if err != nil {
				log.Error().Str("module", "orders").Int64("height", tx.Height).
					Err(err).Msg("error while registering orders")
				return fmt.Errorf("failed to save order update: %w", err)
			}
		}
	}

	return nil
}
