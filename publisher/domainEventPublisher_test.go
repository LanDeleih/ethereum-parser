package publisher_test

import (
	"testing"

	"github.com/landeleih/ethereum-parser/common/types/base"
	"github.com/landeleih/ethereum-parser/common/types/events"
	"github.com/landeleih/ethereum-parser/publisher"
)

func TestNewDomainEventPublisher(t *testing.T) {
	t.Parallel()

	newPublisher := publisher.NewDomainEventPublisher()

	id := "test"
	anotherID := "another_test"

	listener := newFakeDomainEventListener(t, id, "fake_domain_event")

	newPublisher.RegisterListeners(listener)

	event := &FakeDomainEvent{id: id, eventType: "fake_domain_event"}
	anotherEvent := &FakeDomainEvent{id: anotherID, eventType: "another_fake_domain_event"}

	newPublisher.Publish(event, anotherEvent)
}

type FakeDomainEventListener struct {
	test      *testing.T
	id        string
	eventType string
}

func newFakeDomainEventListener(
	test *testing.T,
	id string,
	eventType string,
) events.DomainEventListener {
	return &FakeDomainEventListener{
		test:      test,
		id:        id,
		eventType: eventType,
	}
}

func (c *FakeDomainEventListener) EventType() base.DomainEvent {
	return &FakeDomainEvent{eventType: c.eventType}
}

func (c *FakeDomainEventListener) Handle(event base.DomainEvent) {
	if event.Type() == "" || event.ID() == "" {
		c.test.Log("got an empty event")
		c.test.Fail()
	}
	if event.ID() != c.id {
		c.test.Log("got wrong event id")
		c.test.Fail()
	}
}

type FakeDomainEvent struct {
	id        string
	eventType string
}

func (m *FakeDomainEvent) ID() string {
	return m.id
}

func (m *FakeDomainEvent) Type() string {
	return m.eventType
}
