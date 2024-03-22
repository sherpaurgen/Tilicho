package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// migration will make sure database schema is properly initialized or updated
func (db *Database) Migratedb() error {
	fmt.Println("starting migration of database ..")
	connString := db.ConnectionString // Adjust this line accordingly
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		return err
	}
	sqlDB := stdlib.OpenDB(*config)

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations", "postgres", driver,
	)
	if err != nil {
		fmt.Println("error occured in migrate.NewWithDatabaseInstance: %w", err)
		return err
	}

	if err := m.Up(); err != nil {
		fmt.Printf("could not run m.Up() error : %v\n", err)
		return err
	}
	fmt.Println("successfully migrated...")
	return nil
}
