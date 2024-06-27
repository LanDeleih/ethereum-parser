package domain

type Transaction struct {
	BlockNumber string `json:"block_number,omitempty"`
	BlockHash   string `json:"block_hash,omitempty"`
	Hash        string `json:"hash,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Value       string `json:"value,omitempty"`
}
