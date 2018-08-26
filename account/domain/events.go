package domain

import (
	eh "github.com/looplab/eventhorizon"
)

const (
	// AccountCreatedEvent means an account was created for a user
	AccountCreatedEvent eh.EventType = "AccountCreated"

	// AccountToppedUpEvent means more funds added to an account
	AccountToppedUpEvent eh.EventType = "AccountToppedUp"

	// AccountTipDeductedEvent means funds were used for a tip
	AccountTipDeductedEvent eh.EventType = "AccountTipDeducted"
)

func init() {
	eh.RegisterEventData(AccountToppedUpEvent, func() eh.EventData {
		return &AccountToppedUpData{}
	})
	eh.RegisterEventData(AccountTipDeductedEvent, func() eh.EventData {
		return &AccountTipDeductedData{}
	})
}

// AccountToppedUpData is the event data for AccountToppedUpEvent
type AccountToppedUpData struct {
	TopUpAmount float32 `json:"top_up_amount"`
}

// AccountTipDeductedData is the event data for AccountTipDeductedEvent
type AccountTipDeductedData struct {
	DeductedAmount float32 `json:"deducted_amount"`
}
