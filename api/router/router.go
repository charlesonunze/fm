package router

import (
	"log/slog"

	"fm/api/handlers"
	"fm/db"
	transactionsSvc "fm/services/transactions"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(g *echo.Group, DB *mongo.Database, logger *slog.Logger) {
	// initialize health handler
	hh := handlers.NewHealthHandler()
	// setup health route
	g.GET("/health", hh.CheckHealth)

	// initialize payments handler
	depositsRepo := db.NewDepositsRepo(DB)
	paymentsSvc := transactionsSvc.NewTransactionsService(depositsRepo, logger)
	paymentsHandler := handlers.NewPaymentsHandler(paymentsSvc)
	// setup payment routes
	tpRouter := g.Group("/third-party")
	paymentsRouter := tpRouter.Group("/payments")
	paymentsRouter.POST("/", paymentsHandler.CreateDeposit)
}
