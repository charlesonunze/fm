package integrationtest

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"fm/db"

	"github.com/maragudk/env"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateDatabase for testing.
// Usage:
//
//	db, cleanup := CreateDatabase()
//	defer cleanup()
//	…
func CreateDatabase() (*db.Database, func()) {
	wd, _ := os.Getwd()
	root, _ := GetProjectRoot(wd)
	env.MustLoad(filepath.Join(root, ".env.test"))

	return connect()
}

func connect() (*db.Database, func()) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	db := db.New(db.NewDatabaseOptions{Logger: logger})
	dbName := env.GetStringOrDefault("DB_NAME", "fm_test")
	if err := db.Connect(dbName); err != nil {
		panic(err)
	}

	return db, func() {
		// Get a list of all collections in the database
		collections, err := db.DB.ListCollectionNames(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		// Drop each collection
		for _, collection := range collections {
			err := db.DB.Collection(collection).Drop(context.Background())
			if err != nil {
				log.Printf("Failed to drop collection %s: %v\n", collection, err)
			} else {
				log.Printf("Dropped collection: %s\n", collection)
			}
		}

		if err := db.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}
}

func GetProjectRoot(currentFile string) (string, error) {
	// Walk up the directory hierarchy until a file named "go.mod" is found
	for {
		dir := filepath.Dir(currentFile)
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		if dir == filepath.Dir(dir) { // Reached root directory
			return "", errors.New("project root not found")
		}
		currentFile = dir
	}
}
