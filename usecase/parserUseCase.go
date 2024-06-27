package usecase

import "github.com/landeleih/ethereum-parser/domain"

type ParserUseCase interface {
	// GetCurrentBlock last parsed block
	GetCurrentBlock() int
	// Subscribe add address to observer
	Subscribe(address string) bool
	// GetTransactions list of inbound or outbound transactions for an address
	GetTransactions(address string) []domain.Transaction
}
