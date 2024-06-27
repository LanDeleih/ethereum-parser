package access

import "github.com/landeleih/ethereum-parser/domain"

type SubscriptionPersister interface {
	Subscribe(subscription domain.Subscription) bool
}
