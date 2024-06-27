package access

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type TransactionExtractor interface {
	ByAddress(ctx context.Context, address string) ([]domain.Transaction, error)
}
