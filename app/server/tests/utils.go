package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

var db *sql.DB
var ctx context.Context

type ResponseBody struct {
	Message string `json:"message"`
}

func runMigrations(db *sql.DB, order string) error {

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/scripts",
		"postgres", driver,
	)

	if err != nil {
		return errors.New("could not run migrations")
	}

	switch order {
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		return errors.New("migration run bad argument")
	}

	return nil
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
