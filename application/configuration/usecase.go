package configuration

import (
	"github.com/landeleih/ethereum-parser/usecase"
	"github.com/landeleih/ethereum-parser/usecase/access"
	"github.com/landeleih/ethereum-parser/usecase/provider"
	"github.com/landeleih/ethereum-parser/usecase/scenarios"
)

func parserUseCase(
	blockProvider provider.BlockProvider,
	transactionExtractor access.TransactionExtractor,
	subscribtionPersister access.SubscriptionPersister,
) usecase.ParserUseCase {
	return scenarios.NewParserUseCase(
		blockProvider,
		transactionExtractor,
		subscribtionPersister,
	)
}
