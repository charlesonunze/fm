package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID            primitive.ObjectID `bson:"_id"`
	Type          string             `bson:"type"`
	IdempotencyID string             `bson:"idempotency_id"`
	AccountID     string             `bson:"account_id"`
	Reference     string             `bson:"reference"`
	Amount        float64            `bson:"amount"`
	CreatedAt     time.Time          `bson:"createdAt"`
	Response      map[string]any     `bson:"response"`
}
