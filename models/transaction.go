package models

import "time"

type Transaction struct {
	ID        string          `json:"id"`
	AccountID string          `json:"account_id"`
	Type      TransactionType `json:"type"`
	Reference string          `json:"reference"`
	Amount    float64         `json:"amount"`
	CreatedAt time.Time       `json:"createdAt"`
}

type TransactionType int

const (
	TransactionTypeDeposit TransactionType = iota + 1
	TransactionTypeWithdraw
)

func (t TransactionType) String() string {
	switch t {
	case TransactionTypeDeposit:
		return "deposit"
	case TransactionTypeWithdraw:
		return "withdraw"
	}
	return ""
}
