package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthHandler struct{}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "I'm alive!")
}
