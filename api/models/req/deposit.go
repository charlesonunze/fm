package req

import (
	"time"
)

type CreateDeposit struct {
	ID            string    `json:"id"`
	IdempotencyID string    `json:"idempotency_id"`
	AccountID     string    `json:"account_id"`
	Reference     string    `json:"reference"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `bson:"createdAt"`
}
