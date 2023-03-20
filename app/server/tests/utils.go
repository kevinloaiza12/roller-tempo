package tests

import (
	"database/sql"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type ResponseBody struct {
	Message string `json:"message"`
}

func runMigrations(t *testing.T, db *sql.DB, order string) error {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	failOnError(t, err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/scripts",
		"postgres", driver)
	failOnError(t, err)

	switch order {
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		t.Error("migration run bad argument")
	}

	return nil
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
