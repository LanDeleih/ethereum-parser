package configuration

import (
	"log"
	"net/http"

	"github.com/landeleih/ethereum-parser/application/listener"
	"github.com/landeleih/ethereum-parser/application/scheduler"
	persistence "github.com/landeleih/ethereum-parser/in-memory-storage"
	providerimpl "github.com/landeleih/ethereum-parser/json-rpc"
	"github.com/landeleih/ethereum-parser/publisher"
	"github.com/landeleih/ethereum-parser/usecase/access"
)

type Application interface {
	Run() error
}

type Configuration struct {
	handlers  http.Handler
	scheduler *scheduler.Scheduler
}

func Initialize() Application {
	domainEventPublisher := publisher.NewDomainEventPublisher()

	var subscriberRepository = persistence.NewInMemorySubscriberRepository(domainEventPublisher)
	var jsonrpcRepository = providerimpl.NewJSONRPCRepository("", "", 0)
	var transactionExtractor access.TransactionExtractor = persistence.NewInMemoryTransactionRepository()
	var transactionRepository = persistence.NewInMemoryTransactionRepository()

	newParserUseCase := parserUseCase(
		jsonrpcRepository,
		transactionExtractor,
		subscriberRepository,
	)
	sch := scheduler.NewScheduler()
	subscriberListener := listener.NewSubsciptionCreatedListener(
		subscriberRepository,
		jsonrpcRepository,
		transactionRepository,
		sch,
	)

	domainEventPublisher.RegisterListeners(subscriberListener)

	handlers := newRestHandlers(newParserUseCase)

	return &Configuration{
		handlers:  handlers,
		scheduler: sch,
	}
}

func (c *Configuration) Run() error {
	log.Println("starting scheduler")
	go c.scheduler.Start()

	log.Println("Starting server on port 8080")
	return http.ListenAndServe(":8080", c.handlers)
}
