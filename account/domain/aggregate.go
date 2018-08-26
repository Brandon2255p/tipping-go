package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

// AccountAggregateType is the type name for this aggregate
const AccountAggregateType eh.AggregateType = "AccountAggregate"

var TimeNow = time.Now

func init() {
	eh.RegisterAggregate(func(id eh.UUID) eh.Aggregate {
		return &AccountAggregate{
			AggregateBase: events.NewAggregateBase(AccountAggregateType, id),
		}
	})
}

type AccountAggregate struct {
	*events.AggregateBase
	balance float32
}

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *AccountAggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case *CreateAccount:
		a.StoreEvent(AccountCreatedEvent, nil, TimeNow())
	case *TopUpAccount:
		a.StoreEvent(AccountToppedUpEvent, &AccountToppedUpData{
			TopUpAmount: cmd.Amount,
		}, TimeNow())
	case *PayTipFromAccount:
		negativeBalance := a.balance-cmd.Amount < 0
		if negativeBalance {
			return errors.New("Cant deduct more than the account balance")
		}
		a.StoreEvent(AccountTipDeductedEvent, &AccountTipDeductedData{
			DeductedAmount: cmd.Amount,
		}, TimeNow())
	default:
		return fmt.Errorf("could not handle command: %s", cmd.CommandType())
	}
	return nil
}

// ApplyEvent implements the ApplyEvent method of the
// eventhorizon.Aggregate interface.
func (a *AccountAggregate) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case AccountCreatedEvent:

	case AccountToppedUpEvent:
		data, ok := event.Data().(*AccountToppedUpData)
		if !ok {
			return errors.New("invalid event data")
		}
		a.balance += data.TopUpAmount
	case AccountTipDeductedEvent:
		data, ok := event.Data().(*AccountTipDeductedData)
		if !ok {
			return errors.New("invalid event data")
		}
		a.balance -= data.DeductedAmount
	default:
		return fmt.Errorf("could not apply event: %s", event.EventType())
	}
	return nil
}
