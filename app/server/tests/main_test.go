package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
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

func TestCoinsAccum(t *testing.T) {

	initialCoins := 15000
	rewardedCoins := 18000

	inputUsr := models.NewUser(1193132712, initialCoins, 0)

	_, err = database.CreateNewUser(ctx, db, inputUsr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	inputUsr.SetUserCoins(rewardedCoins)

	_, err = database.UsersUpdateQuery(ctx, db, inputUsr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputUsr, err := database.GetUserByID(ctx, db, 1193132712)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputVal := outputUsr.GetUserCoins()

	if outputVal != rewardedCoins {
		t.Error("Input difers from output")
	}
}

func TestAllAttractionView(t *testing.T) {

	input1 := models.NewAttraction("Ruleta #2", "Es una gran ruleta", 150, 30, 0, 0)
	input2 := models.NewAttraction("Canal del amor", "Un romántico paseo en bote para los más tortolitos", 260, 20, 0, 1)

	_, err = database.CreateNewAttraction(ctx, db, input1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = database.CreateNewAttraction(ctx, db, input2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err := database.GetAllAttractions(ctx, db)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMeanTimeView(t *testing.T) {

	inputMeanTime := 120
	input := models.NewAttraction("Kamikaze", "Atracción que pondrá a prueba tus nervios.", inputMeanTime, 20, 0, 1)

	_, err = database.CreateNewAttraction(ctx, db, input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	output, err := database.GetAttractionByName(ctx, db, "Kamikaze")
	if err != nil {
		t.Fatalf(err.Error())
	}
	outputMeanTime := output.GetAttractionDuration()

	if !reflect.DeepEqual(outputMeanTime, inputMeanTime) {
		t.Error("Input difers from output")
	}

}

func TestGetTurn(t *testing.T) {

	inputAtt := models.NewAttraction("Carritos chocones", "Choca a todos los que puedas.", 150, 30, 12, 13)
	inputUsr := models.NewUser(1193132714, 16900, 0)

	_, err = database.CreateNewAttraction(ctx, db, inputAtt)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = database.CreateNewUser(ctx, db, inputUsr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputAtt, err := database.GetAttractionByName(ctx, db, "Carritos chocones")
	if err != nil {
		t.Fatalf(err.Error())
	}

	currentTurn := outputAtt.GetAttractionCurrentTurn()
	nextTurn := currentTurn + 1

	outputAtt.SetAttractionCurrentTurn(nextTurn)

	inputUsr.SetUserTurn(currentTurn)

	_, err = database.UsersUpdateQuery(ctx, db, inputUsr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputUsr, err := database.GetUserByID(ctx, db, 1193132714)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputCurrentTurn := outputUsr.GetUserTurn()

	if !reflect.DeepEqual(outputUsr, inputUsr) {
		t.Error("Input difers from output")
	}
	if !reflect.DeepEqual(outputCurrentTurn, currentTurn) {
		t.Error("Input difers from output")
	}
}
