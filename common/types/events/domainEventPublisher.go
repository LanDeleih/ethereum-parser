package events

import "github.com/landeleih/ethereum-parser/common/types/base"

type DomainEventPublisher interface {
	Publish(events ...base.DomainEvent)
}
