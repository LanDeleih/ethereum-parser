package domain

import (
	"github.com/landeleih/ethereum-parser/common/types/base"
)

type SubscriptionCreatedDomainEvent struct {
	id string
}

func (m *SubscriptionCreatedDomainEvent) ID() string { return m.id }

func (m *SubscriptionCreatedDomainEvent) Type() string { return "subscription created" }

func NewSubscriptionCreatedDomainEvent(
	id string,
) base.DomainEvent {
	return &SubscriptionCreatedDomainEvent{
		id: id,
	}
}
