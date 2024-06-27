package persistence

import (
	"context"
	"slices"
	"sync"

	"github.com/landeleih/ethereum-parser/domain"
)

type InMemoryTransactionRepository struct {
	transactions map[string][]domain.Transaction
	mytex        sync.RWMutex
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[string][]domain.Transaction),
		mytex:        sync.RWMutex{},
	}
}

func (i *InMemoryTransactionRepository) ByAddress(
	_ context.Context,
	address string,
) ([]domain.Transaction, error) {
	i.mytex.RLock()
	defer i.mytex.RUnlock()

	txs, ok := i.transactions[address]
	if !ok {
		return nil, ErrTransactionNotFound
	}
	return txs, nil
}

func (i *InMemoryTransactionRepository) Save(
	_ context.Context,
	address string,
	transaction domain.Transaction,
) error {
	i.mytex.Lock()
	defer i.mytex.Unlock()

	transactions, ok := i.transactions[address]
	if !ok {
		i.transactions[address] = append(transactions, transaction)
		return nil
	}

	if slices.Contains(transactions, transaction) {
		return ErrTransactionAlreadyExist
	}

	transactions = append(transactions, transaction)
	i.transactions[address] = transactions

	return nil
}
