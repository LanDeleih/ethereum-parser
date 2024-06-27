package listener

import (
	"context"
	"log"
	"time"

	"github.com/landeleih/ethereum-parser/application/scheduler"
	"github.com/landeleih/ethereum-parser/common/types/base"
	"github.com/landeleih/ethereum-parser/domain"
	"github.com/landeleih/ethereum-parser/usecase/access"
	"github.com/landeleih/ethereum-parser/usecase/provider"
)

type SubsciptionCreatedListener struct {
	extractor           access.SubscriptionExtractor
	transactionProvider provider.TransactionProvider
	persister           access.TransactionPersister
	sch                 *scheduler.Scheduler
}

func NewSubsciptionCreatedListener(
	extractor access.SubscriptionExtractor,
	transactionProvider provider.TransactionProvider,
	persister access.TransactionPersister,
	sch *scheduler.Scheduler,
) *SubsciptionCreatedListener {
	return &SubsciptionCreatedListener{
		extractor:           extractor,
		transactionProvider: transactionProvider,
		persister:           persister,
		sch:                 sch,
	}
}

func (c *SubsciptionCreatedListener) EventType() base.DomainEvent {
	return &domain.SubscriptionCreatedDomainEvent{}
}

func (c *SubsciptionCreatedListener) Handle(event base.DomainEvent) {
	subscriptionTask := NewSubscriptionTask(
		event.ID(),
		c.extractor,
		c.persister,
		c.transactionProvider,
	)
	log.Printf("[Listener] subscription task created. id: %s", event.ID())
	c.sch.AddTask(subscriptionTask, 30*time.Second)
}

type SubscriptionTask struct {
	id                  string
	extractor           access.SubscriptionExtractor
	persister           access.TransactionPersister
	transactionProvider provider.TransactionProvider
}

func NewSubscriptionTask(
	id string,
	extractor access.SubscriptionExtractor,
	persister access.TransactionPersister,
	transactionProvider provider.TransactionProvider,
) *SubscriptionTask {
	return &SubscriptionTask{
		id:                  id,
		extractor:           extractor,
		persister:           persister,
		transactionProvider: transactionProvider,
	}
}

func (t *SubscriptionTask) Do() {
	subscription, err := t.extractor.ByID(t.id)
	if err != nil {
		log.Println("Error while extracting subscription", err)
		return
	}

	transactions, err := t.transactionProvider.ByAddress(context.TODO(), subscription.Address)
	if err != nil {
		log.Println("Error while extracting transactions", err)
		return
	}
	for _, transaction := range transactions {
		err = t.persister.Save(context.TODO(), subscription.Address, transaction)
		if err != nil {
			log.Println("Error while saving subscription", err)
			continue
		}
	}
}
