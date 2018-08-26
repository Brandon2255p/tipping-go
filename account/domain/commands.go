package domain

import (
	eh "github.com/looplab/eventhorizon"
)

const (
	// CreateAccountCommand is the command to initiate an account
	CreateAccountCommand eh.CommandType = "CreateAccount"
	// TopUpAccountCommand is the command to add funds to an account
	TopUpAccountCommand eh.CommandType = "TopUpAccount"
	// PayTipFromAccountCommand is the command to deduct money for a tip payment
	PayTipFromAccountCommand eh.CommandType = "PayTipFromAccount"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &CreateAccount{} })
	eh.RegisterCommand(func() eh.Command { return &TopUpAccount{} })
	eh.RegisterCommand(func() eh.Command { return &PayTipFromAccount{} })
}

// CreateAccount the data structure for the command
type CreateAccount struct {
	ID eh.UUID `json:"id"`
}

func (c CreateAccount) AggregateID() eh.UUID            { return c.ID }
func (c CreateAccount) AggregateType() eh.AggregateType { return AccountAggregateType }
func (c CreateAccount) CommandType() eh.CommandType     { return CreateAccountCommand }

// TopUpAccount the data structure for the command
type TopUpAccount struct {
	ID     eh.UUID
	Amount float32
}

func (c TopUpAccount) AggregateID() eh.UUID            { return c.ID }
func (c TopUpAccount) AggregateType() eh.AggregateType { return AccountAggregateType }
func (c TopUpAccount) CommandType() eh.CommandType     { return TopUpAccountCommand }

// PayTipFromAccount the data structure for the command
type PayTipFromAccount struct {
	ID     eh.UUID
	Amount float32
}

func (c PayTipFromAccount) AggregateID() eh.UUID            { return c.ID }
func (c PayTipFromAccount) AggregateType() eh.AggregateType { return AccountAggregateType }
func (c PayTipFromAccount) CommandType() eh.CommandType     { return PayTipFromAccountCommand }

// Static command type checks
var _ = eh.Command(CreateAccount{})
var _ = eh.Command(TopUpAccount{})
var _ = eh.Command(PayTipFromAccount{})
