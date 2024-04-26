// Package main is the entry point to the server. It reads configuration, sets up logging and error handling,
// handles signals from the OS, and starts and stops the server.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fm/db"
	"fm/server"

	"github.com/maragudk/env"
	slogFormatter "github.com/samber/slog-formatter"
	"golang.org/x/sync/errgroup"
)

// release is set through the linker at build time, generally from a git sha.
// Used for logging and error reporting.
var release string

func main() {
	os.Exit(start())
}

func start() int {
	// Load env variables
	_ = env.Load()

	// Create a slog logger, which:
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	logger := slog.New(
		slogFormatter.NewFormatterHandler(
			slogFormatter.TimezoneConverter(time.UTC),
			slogFormatter.TimeFormatter(time.RFC3339, nil),
		)(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
		),
	)

	// Set logger attributes
	logEnv := env.GetStringOrDefault("LOG_ENV", "development")
	logger = logger.With("env", logEnv)
	logger = logger.With("release", release)

	port := fmt.Sprintf(":%d", env.GetIntOrDefault("PORT", 8080))

	s := server.New(server.Options{
		Database: db.New(db.NewDatabaseOptions{Logger: logger}),
		Port:     port,
		Logger:   logger,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.Start(); err != nil {
			logger.Error(fmt.Errorf("error starting server %w", err).Error())
			return err
		}
		return nil
	})

	<-ctx.Done()

	eg.Go(func() error {
		if err := s.Stop(); err != nil {
			logger.Error(fmt.Errorf("error stopping server %w", err).Error())
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 1
	}

	return 0
}
