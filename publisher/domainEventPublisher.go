package publisher

import (
	"log"

	"github.com/landeleih/ethereum-parser/common/types/base"
	"github.com/landeleih/ethereum-parser/common/types/events"
)

type DomainEventPublisher struct {
	listeners map[string][]events.DomainEventListener
}

func NewDomainEventPublisher() *DomainEventPublisher {
	listeners := make(map[string][]events.DomainEventListener)
	return &DomainEventPublisher{
		listeners: listeners,
	}
}

func (d *DomainEventPublisher) RegisterListeners(
	domainEventListeners ...events.DomainEventListener,
) {
	for _, domainEventListener := range domainEventListeners {
		domainEvent := domainEventListener.EventType()
		d.listeners[domainEvent.Type()] = append(d.listeners[domainEvent.Type()], domainEventListener)
	}
}

func (d *DomainEventPublisher) Publish(events ...base.DomainEvent) {
	for _, event := range events {
		listener, ok := d.listeners[event.Type()]
		if !ok {
			log.Println("listener for event not found", "event", event.Type())
			continue
		}
		log.Println("event published", "id", event.ID(), "event", event.Type())
		d.sendEvent(listener, event)
	}
}

func (d *DomainEventPublisher) sendEvent(
	listener []events.DomainEventListener,
	event base.DomainEvent,
) {
	for l := range listener {
		listener[l].Handle(event)
	}
}
