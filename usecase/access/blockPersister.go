package access

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type BlockPersister interface {
	Save(ctx context.Context, block domain.Block) error
}
