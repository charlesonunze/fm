package models

import "time"

type Deposit struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Reference string    `json:"reference"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}
