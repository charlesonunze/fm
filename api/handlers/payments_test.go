package handlers_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"fm/api/handlers"
	"fm/api/models/req"
	repo "fm/db"
	"fm/integrationtest"
	svc "fm/services/transactions"

	"github.com/stretchr/testify/assert"
)

func TestCreateDepositHandler(t *testing.T) {

	db, cleanup := integrationtest.CreateDatabase()
	defer cleanup()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	repo := repo.NewDepositsRepo(db.DB)
	svc := svc.NewTransactionsService(repo, logger)
	paymentsHandler := handlers.NewPaymentsHandler(svc)

	t.Run("creates an payment", func(t *testing.T) {
		deposit := req.CreateDeposit{
			IdempotencyID: "ididid",
			AccountID:     "123",
			Reference:     "123",
			Amount:        12.2,
			CreatedAt:     time.Now().UTC(),
		}
		jsonData, err := json.Marshal(deposit)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		ctx, res := makePostRequest("/third-party/payments", strings.NewReader(string(jsonData)))
		err = paymentsHandler.CreateDeposit(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
	})

}
