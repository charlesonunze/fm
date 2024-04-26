package transform

import (
	dbModel "fm/db/models"
	appModel "fm/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToDbTransaction converts the deposit app model to the matching db model
//
// Usage:
//
//	dbTransaction := transform.ToTransactionDbModel(deposit)
//	…
func ToTransactionDbModel(deposit appModel.Transaction) dbModel.Transaction {
	// if there is no deposit.ID generate a new ObjectID
	// else generate a new ObjectID from deposit.ID
	var depositId primitive.ObjectID
	if deposit.ID == "" {
		depositId = primitive.NewObjectID()
	} else {
		depositId, _ = primitive.ObjectIDFromHex(deposit.ID)
	}
	return dbModel.Transaction{
		ID:        depositId,
		AccountID: deposit.AccountID,
		Reference: deposit.Reference,
		Amount:    deposit.Amount,
		CreatedAt: deposit.CreatedAt,
	}
}

func ToTransactionAppModel(deposit dbModel.Transaction) *appModel.Transaction {
	return &appModel.Transaction{
		ID:        deposit.ID.Hex(),
		AccountID: deposit.AccountID,
		Reference: deposit.Reference,
		Amount:    deposit.Amount,
		CreatedAt: deposit.CreatedAt,
	}
}
