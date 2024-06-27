package provider_test

import (
	"context"
	"testing"
	"time"

	"github.com/landeleih/ethereum-parser/domain"
	provider "github.com/landeleih/ethereum-parser/json-rpc"
)

func TestJSONRPCRepository_LatestBlockSuccess(t *testing.T) {
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
}
