package db

import (
	"context"
	"fmt"

	dbModel "fm/db/models"
	appModel "fm/models"
	"fm/transform"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionsRepo interface {
	CreateTransaction(ctx context.Context, transaction appModel.Transaction) (*dbModel.Transaction, error)
	GetTransactionByIdempotency(ctx context.Context, id string) (*dbModel.Transaction, error)
}

type transactionsRepo struct {
	db   *mongo.Database
	coll *mongo.Collection
}

const transactionColl = "transactions"

func NewTransactionsRepo(db *mongo.Database) TransactionsRepo {
	return &transactionsRepo{db, db.Collection(transactionColl)}
}

// CreateTransaction creates a new transaction in the database.
func (r *transactionsRepo) CreateTransaction(ctx context.Context, transaction appModel.Transaction) (*dbModel.Transaction, error) {
	dbTransaction := transform.ToTransactionDbModel(transaction)
	res, err := r.coll.InsertOne(ctx, dbTransaction)
	if err != nil || res.InsertedID == nil {
		return nil, fmt.Errorf("[db]: error inserting transaction: %w", err)
	}

	return r.GetTransactionByID(ctx, res.InsertedID.(primitive.ObjectID).Hex())
}

// GetTransactionByID retrieves a transaction from the database by its ID.
func (r *transactionsRepo) GetTransactionByID(ctx context.Context, transactionID string) (*dbModel.Transaction, error) {
	var transaction dbModel.Transaction

	// convert id string to ObjectID
	_id, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return nil, fmt.Errorf("[db]: error converting to object id: %w", err)
	}

	if err = r.coll.
		FindOne(ctx, bson.M{"_id": _id}).
		Decode(&transaction); err != nil {
		// return nil if no document is found
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("[db]: error finding document: %w", err)
	}

	return &transaction, nil
}

// GetTransactionByIdempotency retrieves a transaction from the database by its idempotency id.
func (r *transactionsRepo) GetTransactionByIdempotency(ctx context.Context, id string) (*dbModel.Transaction, error) {
	var transaction dbModel.Transaction

	if err := r.coll.
		FindOne(ctx, bson.M{"idempotency_id": id}).
		Decode(&transaction); err != nil {
		// return nil if no document is found
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("[db]: error finding document: %w", err)
	}

	return &transaction, nil
}

// GetTransactionByReference retrieves a transaction from the database by its name.
func (r *transactionsRepo) GetTransactionByReference(ctx context.Context, reference string) (*dbModel.Transaction, error) {
	var transaction dbModel.Transaction

	if err := r.coll.
		FindOne(ctx, bson.M{"name": reference}).
		Decode(&transaction); err != nil {
		// return nil if no document is found
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("[db]: error finding document: %w", err)
	}

	return &transaction, nil
}
