package marginaccwithdraw

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	log.Info().Str("module", "marginaccwithdraw").
		Int64("height", tx.Height).
		Str("tx", tx.TxHash).
		Msg("HandleTx")
	// Iterate through the events in the transaction directly
	for _, event := range tx.Events {
		// Check if the event type matches the order events we're interested in
		if event.Type == types.EventTypeMsgWithdraw {
			// Initialize a map to store the attributes of the event
			attributes := make(map[string]string)
			for _, attr := range event.Attributes {
				attributes[attr.Key] = attr.Value
			}

			// Extract the attributes we're interested in
			marginAccAddr := attributes[types.AttributeKeyMarginAccountAddress]
			recipient := attributes[types.AttributeKeyRecipiant]
			withdrawCoin := attributes[types.AttributeKeyAmount]
			// Save the order event using the extracted attributes
			err := m.db.SaveWithdrawEvent(types.NewWithdrawEvent(
				marginAccAddr,
				recipient,
				withdrawCoin,
			))
			if err != nil {
				log.Error().Str("module", "marginaccwithdraw").Int64("height", tx.Height).
					Err(err).Msg("error while registering marginaccwithdraw")
				return fmt.Errorf("failed to save marginaccwithdraw event: %w", err)
			}
		}
	}

	return nil
}
