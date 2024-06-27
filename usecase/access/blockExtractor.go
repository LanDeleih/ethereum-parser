package access

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type BlockExtractor interface {
	LatestBlock(ctx context.Context, id int, parameters ...any) (domain.Block, error)
}
