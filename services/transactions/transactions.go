package transactions

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	reqModel "fm/api/models/req"
	"fm/db"
	appModel "fm/models"
	"fm/services"
	"fm/services/tp"
)

type TransactionsService interface {
	Deposit(ctx context.Context, req reqModel.CreateDeposit) (*appModel.Deposit, error)
}

type paymentsService struct {
	name      string
	repo      db.TransactionsRepo
	logger    *slog.Logger
	tpService tp.TpService
}

func NewTransactionsService(repo db.TransactionsRepo, logger *slog.Logger, tpService tp.TpService) TransactionsService {
	return &paymentsService{"payments_service", repo, logger, tpService}
}

func (s *paymentsService) Deposit(ctx context.Context, deposit reqModel.CreateDeposit) (*appModel.Deposit, error) {
	// check if an deposit with the same idempotency id already exists
	d, err := s.repo.GetTransactionByIdempotency(ctx, deposit.IdempotencyID)
	if err != nil {
		svcErr := services.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Errorf("error getting deposit with idempotency id [%v]: %w", deposit.IdempotencyID, err).Error(),
		}
		s.logger.Error(svcErr.LogMsg(s.name))
		return nil, svcErr
	}

	// check if deposit exists
	if d != nil {
		svcErr := services.ServiceError{
			Code: http.StatusBadRequest,
			Msg:  "duplicate transaction",
		}
		s.logger.Error(svcErr.LogMsg(s.name))
		return nil, svcErr
	}

	// deposits
	// make the thirdparty call
	// transactions
	// update user balance

	_, err = s.repo.CreateTransaction(ctx, appModel.Transaction{
		AccountID: deposit.AccountID,
		Reference: deposit.Reference,
		Amount:    deposit.Amount,
		CreatedAt: deposit.CreatedAt,
	})
	if err != nil {
		svcErr := services.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Errorf("error creating deposit with ref [%v]: %w", deposit.Reference, err).Error(),
		}
		s.logger.Error(svcErr.LogMsg(s.name))
		return nil, svcErr
	}

	result, err := s.tpService.CreateDeposit()
	if err != nil {
		// rollback txn
		return nil, err
	}
	if result.Status == "failed" {
		// rollback txn
	}

	return nil, nil
	// return transform.ToTransactionAppModel(*dbDeposit), nil
}
