package account

import (
	"log"
	"net/http"

	"github.com/Brandon2255p/tipping-go/account/domain"
	"github.com/Brandon2255p/tipping-go/account/readmodel"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/memory"
	"github.com/looplab/eventhorizon/httputils"
	eventpublisher "github.com/looplab/eventhorizon/publisher/local"
)

// Setup will link all the required modules togther
func Setup() {
	eventBus := eventbus.NewEventBus()
	commandBus := bus.NewCommandHandler()
	publisher := eventpublisher.NewEventPublisher()
	publisher.AddObserver(&readmodel.Logger{})

	eventBus.AddHandler(eh.MatchAny(), publisher)
	eventStore := eventstore.NewEventStore()
	aggregateStore, err := events.NewAggregateStore(eventStore, eventBus)
	if err != nil {
		log.Fatalf("Could not create aggregate store: %s", err)
	}
	commandHandler, err := aggregate.NewCommandHandler(domain.AccountAggregateType, aggregateStore)
	if err != nil {
		log.Fatalf("Could not create command handler: %s", err)
	}
	commandBus.SetHandler(commandHandler, domain.CreateAccountCommand)
	commandBus.SetHandler(commandHandler, domain.TopUpAccountCommand)
	commandBus.SetHandler(commandHandler, domain.PayTipFromAccountCommand)

	h := http.NewServeMux()
	h.Handle("/api/account/create", httputils.CommandHandler(commandHandler, domain.CreateAccountCommand))
	h.Handle("/api/account/topup", httputils.CommandHandler(commandHandler, domain.TopUpAccountCommand))
	h.Handle("/api/account/deduct", httputils.CommandHandler(commandHandler, domain.PayTipFromAccountCommand))
	http.ListenAndServe(":8080", h)
}
