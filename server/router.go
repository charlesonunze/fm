package server

import (
	"log/slog"
	"net/http"

	v1Router "fm/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogEcho "github.com/samber/slog-echo"
)

func (s *Server) setupRoutes() {
	// register middleware
	useMiddlewareAllRequests(s.echo, s.logger)
	// register v1 endpoints
	v1Router.New(s.echo.Group("/api"), s.database.DB, s.logger)
}

func useMiddlewareAllRequests(e *echo.Echo, logger *slog.Logger) {
	// Echo Middleware
	e.Use(slogEcho.NewWithConfig(logger, slogEcho.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
		WithRequestID:    false,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	e.Use(middleware.BodyLimit("2M"))
}
