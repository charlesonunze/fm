package handlers

import (
	"net/http"

	reqModel "fm/api/models/req"
	svc "fm/services/transactions"

	"github.com/labstack/echo/v4"
)

type paymentsHandler struct {
	svc svc.TransactionsService
}

func NewPaymentsHandler(svc svc.TransactionsService) *paymentsHandler {
	return &paymentsHandler{svc}
}

func (h *paymentsHandler) CreateDeposit(c echo.Context) error {
	// read request body into payment model
	var reqBody reqModel.CreateDeposit
	if err := c.Bind(&reqBody); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "error parsing request body",
		}
	}

	// TODO: validate request body

	result, err := h.svc.Deposit(c.Request().Context(), reqBody)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusCreated, result)
}
