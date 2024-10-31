package positions

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
)

func (m *Module) HandleTx(tx *juno.Tx) error {
	// Iterate through the events in the transaction directly
	for _, event := range tx.Events {
		// Check if the event type matches the order events we're interested in
		if event.Type == types.EventTypeMsgCreateMarginAccount {
			// Initialize a map to store the attributes of the event
			attributes := make(map[string]string)
			for _, attr := range event.Attributes {
				attributes[attr.Key] = attr.Value
			}

			// Extract the attributes we're interested in
			marginAccAddr := attributes[types.AttributeKeyMarginAccountAddress]
			owner := attributes[types.AttributeKeyOwner]
			accNumber := attributes[types.AttributeKeyAccountNumber]
			// Save the order event using the extracted attributes
			err := m.db.SaveCreateMarginAccountEvent(types.NewMarginAccountEvent(
				marginAccAddr,
				owner,
				accNumber,
			))
			if err != nil {
				return fmt.Errorf("failed to save margin account event: %w", err)
			}
		}
	}

	return nil
}
