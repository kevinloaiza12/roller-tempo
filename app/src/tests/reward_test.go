package tests

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func TestReward(t *testing.T) {

	input := models.NewReward(1, "Peluche", "Es un lindo peluche", 1235)

	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("postgres", connStr)
	failOnError(t, err)
	defer db.Close()

	failOnError(t, runMigrations(t, db, "down"))
	failOnError(t, runMigrations(t, db, "up"))

	_, err = database.CreateNewReward(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetRewardByID(ctx, db, 1)
	failOnError(t, err)

	if !reflect.DeepEqual(output, input) {
		t.Error("input difers from output")
	}
}
