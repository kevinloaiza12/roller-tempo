package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
)

var db *sql.DB
var ctx context.Context
var err error

type ResponseBody struct {
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	ctx = context.Background()
	envErr := godotenv.Load("../config.env")
	if envErr != nil {
		os.Exit(1)
	}

	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	runMigrations(db, "down")
	runMigrations(db, "up")

	code := m.Run()

	os.Exit(code)
}

func runMigrations(db *sql.DB, order string) error {
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/scripts",
		"postgres", driver,
	)

	if err != nil {
		return errors.New("Could not run migrations")
	}

	switch order {
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		return errors.New("Migration run bad argument")
	}

	return nil
}
