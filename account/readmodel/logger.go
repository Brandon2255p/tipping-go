package readmodel

import (
	"context"
	"log"

	eh "github.com/looplab/eventhorizon"
)

// Logger is a simple event observer for logging all events.
type Logger struct{}

// Notify implements the Notify method of the EventObserver interface.
func (l *Logger) Notify(ctx context.Context, event eh.Event) {
	log.Println("PUBLISHER:", event)
}
