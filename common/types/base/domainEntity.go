package base

type DomainEntity interface {
	AddEvent(event DomainEvent)
	Events() []DomainEvent
}

type DomainEntityImpl struct {
	events []DomainEvent
}

func NewDomainEntity() *DomainEntityImpl {
	events := make([]DomainEvent, 0)
	return &DomainEntityImpl{events: events}
}

func (d *DomainEntityImpl) AddEvent(event DomainEvent) {
	d.events = append(d.events, event)
}

func (d *DomainEntityImpl) Events() []DomainEvent {
	return d.events
}
