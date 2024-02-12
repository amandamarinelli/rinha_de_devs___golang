package models

type Transaction struct {
	ClientId    uint64
	Value       uint64
	Type        TransactionType
	Description string
}

type TransactionType string

type TransactionResponse struct {
	Limit   uint64
	Balance int64
}
