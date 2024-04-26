package integrationtest

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	"fm/db"
	"fm/server"

	"github.com/maragudk/env"
)

// CreateServer for testing on port 7777, returning a cleanup function that stops the server.
// Usage:
//
//	cleanup := CreateServer()
//	defer cleanup()
func CreateServer() func() {
	_ = env.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	s := server.New(server.Options{
		Database: createDatabase(logger),
		Logger:   logger,
	})

	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	for {
		_, err := http.Get("http://localhost:7777/health")
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}

	return func() {
		if err := s.Stop(); err != nil {
			panic(err)
		}
	}
}

// SkipIfShort skips t if the "-short" flag is passed to "go test".
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}

func createDatabase(logger *slog.Logger) *db.Database {
	return db.New(db.NewDatabaseOptions{Logger: logger})
}
