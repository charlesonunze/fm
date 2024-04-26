package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deposit struct {
	ID            primitive.ObjectID `bson:"_id"`
	IdempotencyID string             `bson:"idempotency_id"`
	AccountID     string             `bson:"account_id"`
	Reference     string             `bson:"reference"`
	Amount        float64            `bson:"amount"`
	CreatedAt     time.Time          `bson:"createdAt"`
}
