package domain

type Subscription struct {
	ID      string
	Address string
}

func NewSubscription(
	address string,
) Subscription {
	id := NewID()
	return Subscription{
		ID:      id,
		Address: address,
	}
}
