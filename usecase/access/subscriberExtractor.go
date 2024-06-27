package access

import "github.com/landeleih/ethereum-parser/domain"

type SubscriptionExtractor interface {
	ByID(id string) (domain.Subscription, error)
}
