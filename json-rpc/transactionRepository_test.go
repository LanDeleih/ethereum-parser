package provider_test

import (
	"context"
	"testing"
	"time"

	"github.com/landeleih/ethereum-parser/domain"
	provider "github.com/landeleih/ethereum-parser/json-rpc"
)

func TestJSONRPCRepositoryTransactionsByBlockNumberSuccess(t *testing.T) {
	t.Parallel()

	id := domain.NewID()
	repo := provider.NewJSONRPCRepository("", "", 10*time.Second)
	block, err := repo.LatestBlock(context.TODO(), id)
	if err != nil {
		t.Log("unable to retrieve latest block:", err)
		t.Fail()
	}
	if block.ID != id {
		t.Log("block id does not match")
		t.Fail()
	}

	blockWithTransactions, err := repo.ByBlockNumber(context.TODO(), block.Number, block.ID)
	if err != nil {
		t.Log("unable to retrieve block with transactions:", err)
		t.Fail()
	}
	if len(blockWithTransactions.Transactions) == 0 {
		t.Log("block with transactions is empty")
		t.Fail()
	}
}
