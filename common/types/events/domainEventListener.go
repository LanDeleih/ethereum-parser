package events

import "github.com/landeleih/ethereum-parser/common/types/base"

type DomainEventListener interface {
	EventType() base.DomainEvent
	Handle(event base.DomainEvent)
}
