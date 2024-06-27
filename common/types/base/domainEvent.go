package base

type DomainEvent interface {
	ID() string
	Type() string
}
