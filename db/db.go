package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/maragudk/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is the relational storage abstrpayment.
type Database struct {
	DB     *mongo.Database
	Client *mongo.Client
	logger *slog.Logger
}

// NewDatabaseOptions for NewDatabase.
type NewDatabaseOptions struct {
	Logger *slog.Logger
}

// New with the given options.
// If no logger is provided, logs are discarded.
func New(opts NewDatabaseOptions) *Database {
	if opts.Logger == nil {
		// opts.Logger = zap.NewNop()
	}
	return &Database{
		logger: opts.Logger,
	}
}

// Connect to the database.
func (d *Database) Connect(dbName string) error {
	d.logger.Info("Connecting to database")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dbURI := env.GetStringOrDefault("DB_URI", "mongodb+srv://charles:3W11CmTgeW9jnzIG@cluster0.irtitkv.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	// Set the client and database
	d.Client = client
	d.DB = client.Database(dbName)

	return nil
}

// Close the database.
func (d *Database) Close(ctx context.Context) error {
	if err := d.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

// Ping the database.
func (d *Database) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	var result bson.M
	err := d.Client.Database("admin").
		RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).
		Decode(&result)
	if err != nil {
		return err
	}

	d.logger.Info("Pinged your deployment. You successfully connected to MongoDB!")

	return nil
}
