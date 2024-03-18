package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// pgx pgxpool is used https://github.com/jackc/pgx here because the lib pq is in archive mode and recommends new projects to use pgx
type Database struct {
	Pool             *pgxpool.Pool
	ConnectionString string
}

func NewDatabase() (*Database, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//"postgres://postgres:changeme567@db:5432/postgres?sslmode=disable"
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &Database{Pool: pool, ConnectionString: connectionString}, nil
}

func (db *Database) Ping(ctx context.Context) error {
	//PingContext verifies a connection to the database is still alive, establishing a connection if necessary.
	// basically for healthcheck
	err := db.Pool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}
