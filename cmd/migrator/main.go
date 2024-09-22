package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	dsn := os.Getenv("POSTGRES_DSN")
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Use the directory name as the database name, should be the name of the context.
	dbName := filepath.Base(migrationsPath)

	migrator, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, dbName, driver)
	if err != nil {
		return fmt.Errorf("failed creating migrator: %w", err)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
