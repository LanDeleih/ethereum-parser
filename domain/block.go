package domain

type Block struct {
	ID           string
	Number       int
	Transactions []Transaction
}
