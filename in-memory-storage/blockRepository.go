package persistence

import (
	"context"
	"sync"

	"github.com/landeleih/ethereum-parser/domain"
)

type InMemoryBlockRepository struct {
	blocks map[string]domain.Block
	mytex  sync.RWMutex
}

func (i *InMemoryBlockRepository) Save(_ context.Context, block domain.Block) error {
	i.mytex.Lock()
	defer i.mytex.Unlock()

	i.blocks[block.ID] = block

	return nil
}
