package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	Ctx = context.Background()
	envErr := godotenv.Load("../config.env")
	if envErr != nil {
		os.Exit(1)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		os.Exit(1)
	}

	defer Db.Close()

	runMigrations(Db, "down")
	runMigrations(Db, "up")

	log.Println("Do stuff BEFORE the tests!")
	code := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(code)
}
