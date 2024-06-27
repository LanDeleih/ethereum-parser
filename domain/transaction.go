package domain

type Transaction struct {
	From  string `json:"from,omitempty"`
	To    string `json:"to,omitempty"`
	Value string `json:"value,omitempty"`
}
