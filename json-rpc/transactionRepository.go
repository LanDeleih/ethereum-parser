package provider

import (
	"context"

	"github.com/landeleih/ethereum-parser/domain"
)

func (r *JSONRPCRepository) ByAddress(ctx context.Context, address string) ([]domain.Transaction, error) {
	block, err := r.LatestBlock(ctx, domain.NewID())
	if err != nil {
		return nil, err
	}
	blockWithTransactions, err := r.ByBlockNumber(ctx, block.Number, block.ID)
	if err != nil {
		return nil, err
	}

	transactions := make([]domain.Transaction, 0, len(blockWithTransactions.Transactions))

	for _, transaction := range blockWithTransactions.Transactions {
		if transaction.To == address || transaction.From == address {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}
