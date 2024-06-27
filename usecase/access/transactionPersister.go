package access

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type TransactionPersister interface {
	Save(ctx context.Context, address string, transaction domain.Transaction) error
}
