package db_test

import (
	"context"
	"testing"
	"time"

	repo "fm/db"
	"fm/integrationtest"
	appModel "fm/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDatabase_CreateDeposit(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("creates a new deposit", func(t *testing.T) {
		db, cleanup := integrationtest.CreateDatabase()
		defer cleanup()

		repo := repo.NewDepositsRepo(db.DB)

		arg := makeNewDeposit(t)
		expectedDeposit, err := repo.CreateDeposit(context.Background(), arg)
		assert.NoError(t, err)
		assert.Equal(t, arg.AccountID, expectedDeposit.AccountID)
		assert.Equal(t, arg.Reference, expectedDeposit.Reference)
		assert.Equal(t, arg.Amount, expectedDeposit.Amount)
	})
}

func makeNewDeposit(t *testing.T) appModel.Deposit {
	t.Helper()
	return appModel.Deposit{
		AccountID: primitive.NewObjectID().Hex(),
		Reference: gofakeit.AchAccount(),
		Amount:    gofakeit.Float64(),
		CreatedAt: time.Now().UTC(),
	}
}
