package provider

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type TransactionProvider interface {
	ByAddress(ctx context.Context, address string) ([]domain.Transaction, error)
}
