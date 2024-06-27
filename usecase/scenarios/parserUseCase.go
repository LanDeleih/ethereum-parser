package scenarios

import (
	"context"
	"log"

	"github.com/landeleih/ethereum-parser/domain"
	"github.com/landeleih/ethereum-parser/usecase/access"
	"github.com/landeleih/ethereum-parser/usecase/provider"
)

type ParserUseCase struct {
	blockProvider         provider.BlockProvider
	transactionExtractor  access.TransactionExtractor
	subscriptionPersister access.SubscriptionPersister
}

func NewParserUseCase(
	blockProvider provider.BlockProvider,
	transactionExtractor access.TransactionExtractor,
	subscriptionPersister access.SubscriptionPersister,
) *ParserUseCase {
	return &ParserUseCase{
		blockProvider:         blockProvider,
		transactionExtractor:  transactionExtractor,
		subscriptionPersister: subscriptionPersister,
	}
}

func (p *ParserUseCase) GetCurrentBlock() int {
	id := domain.NewID()
	block, err := p.blockProvider.LatestBlock(context.TODO(), id)
	if err != nil {
		log.Println("[ParserUseCase] error getting latest block", err)
		return -1
	}
	return block.Number
}

func (p *ParserUseCase) Subscribe(address string) bool {
	subscription := domain.NewSubscription(address)
	return p.subscriptionPersister.Subscribe(subscription)
}

func (p *ParserUseCase) GetTransactions(address string) []domain.Transaction {
	transactions, err := p.transactionExtractor.ByAddress(context.TODO(), address)
	if err != nil {
		log.Println("[ParserUseCase] transaction extractor error:", err)
		return nil
	}
	return transactions
}
