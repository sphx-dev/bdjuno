package positions

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	log.Info().Str("module", "positions").
		Int64("height", tx.Height).
		Str("tx", tx.TxHash).
		Msg("HandleTx")
	// Iterate through the events in the transaction directly
	for _, event := range tx.Events {
		// Check if the event type matches the order events we're interested in
		if event.Type == types.EventTypeModifyPosition || event.Type == types.EventTypeNewPosition {
			// Initialize a map to store the attributes of the event
			attributes := make(map[string]string)
			for _, attr := range event.Attributes {
				attributes[attr.Key] = attr.Value
			}

			// Extract the attributes we're interested in
			positionID := attributes[types.AttributeKeyPositionID]
			marginAccountAddress := attributes[types.AttributeKeyMarginAccountAddress]
			marketID := attributes[types.AttributeKeyMarketID]
			positionSize := attributes[types.AttributeKeyPositionSize]
			entryPrice := attributes[types.AttributeKeyEntryPrice]
			leverage := attributes[types.AttributeKeyLeverage]
			entryTime := attributes[types.AttributeKeyEntryTime]
			positionSide := attributes[types.AttributeKeyPositionSide]
			tpOrderID := attributes[types.AttributeKeyTpOrderID]
			slOrderID := attributes[types.AttributeKeySlOrderID]
			positionStatus := attributes[types.AttributeKeyPositionStatus]

			// Save the order event using the extracted attributes
			err := m.db.SavePositionUpdate(types.NewPositionEvent(
				positionID,
				marginAccountAddress,
				marketID,
				positionSize,
				entryPrice,
				leverage,
				entryTime,
				positionSide,
				tpOrderID,
				slOrderID,
				positionStatus,
			))
			if err != nil {
				log.Error().Str("module", "positions").Int64("height", tx.Height).
					Err(err).Msg("error while registering positions")
				return fmt.Errorf("failed to save positions update: %w", err)
			}
		}
	}

	return nil
}
