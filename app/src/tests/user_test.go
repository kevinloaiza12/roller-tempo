package tests

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/resources"
)

func TestUser(t *testing.T) {

	input := resources.NewUser(1193132710, 15000, 0)

	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("postgres", connStr)
	failOnError(t, err)
	defer db.Close()

	failOnError(t, runMigrations(t, db, "down"))
	failOnError(t, runMigrations(t, db, "up"))

	_, err = database.CreateNewUser(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetUserByID(ctx, db, 1193132710)
	failOnError(t, err)

	if !reflect.DeepEqual(output, input) {
		t.Error("input difers from output")
	}
}
