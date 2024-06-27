package persistence

import (
	"sync"

	"github.com/landeleih/ethereum-parser/common/types/events"
	"github.com/landeleih/ethereum-parser/domain"
)

type InMemorySubscriberRepository struct {
	subscribers map[string]string
	mytex       sync.RWMutex
	publisher   events.DomainEventPublisher
}

func NewInMemorySubscriberRepository(
	publisher events.DomainEventPublisher,
) *InMemorySubscriberRepository {
	return &InMemorySubscriberRepository{
		subscribers: make(map[string]string),
		mytex:       sync.RWMutex{},
		publisher:   publisher,
	}
}

func (i *InMemorySubscriberRepository) Subscribe(subscription domain.Subscription) bool {
	i.mytex.Lock()
	defer i.mytex.Unlock()

	i.subscribers[subscription.ID] = subscription.Address

	i.publisher.Publish(domain.NewSubscriptionCreatedDomainEvent(subscription.ID))

	return true
}

func (i *InMemorySubscriberRepository) ByID(id string) (domain.Subscription, error) {
	i.mytex.RLock()
	defer i.mytex.RUnlock()

	address, found := i.subscribers[id]
	if !found {
		return domain.Subscription{}, ErrSubscriptionNotFound
	}

	return domain.Subscription{
		ID:      id,
		Address: address,
	}, nil
}
