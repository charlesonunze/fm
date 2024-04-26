// Package server contains everything for setting up and running the HTTP server.
package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"fm/db"

	"github.com/labstack/echo/v4"
	"github.com/maragudk/env"
)

type Server struct {
	database *db.Database
	logger   *slog.Logger
	echo     *echo.Echo
	port     string
}

type Options struct {
	Database *db.Database
	Logger   *slog.Logger
	Port     string
}

func New(opts Options) *Server {
	return &Server{
		database: opts.Database,
		logger:   opts.Logger,
		echo:     echo.New(),
		port:     opts.Port,
	}
}

func (s *Server) Start() error {
	dbName := env.GetStringOrDefault("DB_NAME", "fm")
	if err := s.database.Connect(dbName); err != nil {
		connectErr := fmt.Errorf("error connecting to database: %w", err)
		s.logger.Error(connectErr.Error())
		return connectErr
	}

	// setup services and handlers
	s.setupRoutes()

	if err := s.echo.Start(s.port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		startErr := fmt.Errorf("error starting server: %w", err)
		s.logger.Error(startErr.Error())
		return startErr
	}

	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.logger.Info("closing the database")

	if err := s.database.Close(ctx); err != nil {
		closeErr := fmt.Errorf("error closing database: %w", err)
		s.logger.Error(closeErr.Error())
		return closeErr
	}

	s.logger.Info("server stopping")

	if err := s.echo.Shutdown(ctx); err != nil {
		shutdownErr := fmt.Errorf("error stopping server: %w", err)
		s.logger.Error(shutdownErr.Error())
		return shutdownErr
	}

	return nil
}
