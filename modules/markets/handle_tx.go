package markets

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	log.Info().Str("module", "markets").
		Int64("height", tx.Height).
		Str("tx", tx.TxHash).
		Msg("HandleTx")
	// Iterate through the events in the transaction directly
	for _, event := range tx.Events {
		// Check if the event type matches the order events we're interested in
		if event.Type == types.EventTypeMsgRegisterMarket {
			// Initialize a map to store the attributes of the event
			attributes := make(map[string]string)
			for _, attr := range event.Attributes {
				attributes[attr.Key] = attr.Value
			}

			// Extract the attributes we're interested in
			marketID := attributes[types.AttributeKeyMarketID]
			marketTicker := attributes[types.AttributeKeyMarketTicker]
			marketBaseAsset := attributes[types.AttributeKeyMarketBaseAsset]
			marketQuoteAsset := attributes[types.AttributeKeyMarketQuoteAsset]
			marketStatus := attributes[types.AttributeKeyMarketStatus]
			// Save the order event using the extracted attributes
			err := m.db.SaveRegisterMarketEvent(types.NewMarketEvent(
				marketID,
				marketTicker,
				marketBaseAsset,
				marketQuoteAsset,
				marketStatus,
			))
			if err != nil {
				log.Error().Str("module", "markets").Int64("height", tx.Height).
					Err(err).Msg("error while registering markets")
				return fmt.Errorf("failed to save markets update: %w", err)
			}
		}
	}

	return nil
}
