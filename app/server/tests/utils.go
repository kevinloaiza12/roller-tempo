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

func TestMain(m *testing.M) {
	envErr := godotenv.Load("../config.env")

	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	runMigrations(db, "down")
	runMigrations(db, "up")

	exitCode := m.Run()

	os.Exit(exitCode)
}

func runMigrations(db *sql.DB, order string) error {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/scripts",
		"postgres", driver)

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
