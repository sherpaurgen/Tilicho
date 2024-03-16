package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// pgx pgxpool is used https://github.com/jackc/pgx here because the lib pq is in archive mode and recommends new projject to use pgx
type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"postgres://user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
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

	return &Database{Pool: pool}, nil
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

type Apartment struct {
	apartment_id   string
	apartment_name string
	address        string
	city           string
	state          string
	zip_code       string
}

type Unit struct {
	unit_id        string
	apartment_id   string
	unit_number    string
	unit_type      string
	square_footage string
	rent           float64
	status         string
	tenant_id      string
}

var (
	ErrFetchingApartment = errors.New("failed to fetch apartment by id")
	ErrNotImplemented    = errors.New("not implemented yet")
)

// all logic will be built on top of Service struct
type Service struct {
	Store Store
}

type Store interface {
	GetApartment(context.Context, string) (Apartment, error)
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetApartment(ctx context.Context, id string) (Apartment, error) {
	apartment, err := s.Store.GetApartment(ctx, id)
	if err != nil {
		return Apartment{}, err
	}
	return apartment, ErrFetchingApartment
}

func (s *Service) UpdateApartment(ctx context.Context, apt Apartment) error {
	//
	return ErrNotImplemented
}
