package provider

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

type BlockProvider interface {
	LatestBlock(ctx context.Context, id string) (domain.Block, error)
}
