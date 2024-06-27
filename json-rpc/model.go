package provider

import "github.com/landeleih/ethereum-parser/domain"

type ModelRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
	Params  []any  `json:"params"`
}

type BlockModelResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      string `json:"id"`
}

type TransactionModelResponse struct {
	JSONRPC string                 `json:"jsonrpc"`
	Result  TransactionModelResult `json:"result"`
	ID      string                 `json:"id"`
}

type TransactionModelResult struct {
	Transactions []TransactionModel `json:"transactions,omitempty"`
}

type TransactionModel struct {
	From  string `json:"from,omitempty"`
	To    string `json:"to,omitempty"`
	Value string `json:"value,omitempty"`
}

func (r BlockModelResponse) toDomainBlock() (domain.Block, error) {
	number, err := ConvertHexToNumber(r.Result)
	if err != nil {
		return domain.Block{}, err
	}
	return domain.Block{
		ID:     r.ID,
		Number: number,
	}, nil
}

func (r TransactionModel) toDomain() domain.Transaction {
	return domain.Transaction{
		From:  r.From,
		To:    r.To,
		Value: r.Value,
	}
}
