package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/landeleih/ethereum-parser/domain"
)

const defaultEthereumURL = "https://cloudflare-eth.com"

const defaultJSONRPCVersion = "2.0"

const defaultTimeout = 15 * time.Second

const (
	EthBlockNumberMethod      = "eth_blockNumber"
	EthGetBlockByNumberMethod = "eth_getBlockByNumber"
)

type JSONRPCRepository struct {
	client         *http.Client
	url            string
	jsonRPCVersion string
}

func NewJSONRPCRepository(
	url string,
	jsonRPCVersion string,
	timeout time.Duration,
) *JSONRPCRepository {
	if url == "" {
		url = defaultEthereumURL
	}
	if jsonRPCVersion == "" {
		jsonRPCVersion = defaultJSONRPCVersion
	}
	if timeout == 0 {
		timeout = defaultTimeout
	}

	client := &http.Client{Timeout: timeout}
	return &JSONRPCRepository{
		client:         client,
		url:            url,
		jsonRPCVersion: jsonRPCVersion,
	}
}

func (r *JSONRPCRepository) LatestBlock(
	ctx context.Context, id string,
) (domain.Block, error) {
	modelRequest, err := r.modelRequest(id, EthBlockNumberMethod)
	if err != nil {
		return domain.Block{}, err
	}
	var modelResponse BlockModelResponse

	err = r.doRequest(modelRequest, &modelResponse)
	if err != nil {
		return domain.Block{}, err
	}
	return modelResponse.toDomainBlock()
}

func (r *JSONRPCRepository) ByBlockNumber(
	_ context.Context,
	blockNumber int,
	id string,
) (domain.Block, error) {
	hexNumber := ConvertNumberToHex(blockNumber)

	modelRequest, err := r.modelRequest(id, EthGetBlockByNumberMethod, hexNumber, true)
	if err != nil {
		return domain.Block{}, err
	}
	var transactionModel TransactionModelResponse
	err = r.doRequest(modelRequest, &transactionModel)
	if err != nil {
		return domain.Block{}, err
	}

	transactions := make([]domain.Transaction, 0, len(transactionModel.Result.Transactions))
	for _, transaction := range transactionModel.Result.Transactions {
		transactions = append(transactions, transaction.toDomain())
	}

	return domain.Block{
		ID:           id,
		Number:       blockNumber,
		Transactions: transactions,
	}, err
}

func (r *JSONRPCRepository) modelRequest(
	id string,
	method string,
	parameters ...any,
) ([]byte, error) {
	return json.Marshal(ModelRequest{
		JSONRPC: r.jsonRPCVersion,
		Method:  method,
		ID:      id,
		Params:  parameters,
	})
}

func (r *JSONRPCRepository) doRequest(requestBody []byte, response any) error {
	resBody, err := r.postRequest(requestBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}

	return nil
}

func (r *JSONRPCRepository) postRequest(reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, r.url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return body, err
	}

	return body, nil
}
