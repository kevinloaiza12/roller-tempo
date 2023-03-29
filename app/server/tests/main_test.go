package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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

	code := m.Run()

	cleanTables(db)

	defer db.Close()

	runMigrations(db, "down")
	runMigrations(db, "up")

	os.Exit(code)
}

func cleanTables(db *sql.DB) {
	db.Exec("DELETE FROM usuarios")
	db.Exec("DELETE FROM atracciones")
	db.Exec("DELETE FROM premios")
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

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCoinsView(t *testing.T) {

	inputVal := 15000
	log.Println("Saldo ingresado: ", inputVal)
	input := models.NewUser(1193132710, inputVal, 0)

	_, err = database.CreateNewUser(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetUserByID(ctx, db, 1193132710)
	failOnError(t, err)

	outputVal := output.GetUserCoins()
	log.Println("Saldo servidor: ", outputVal)

	if !reflect.DeepEqual(outputVal, inputVal) {
		t.Error("Input difers from output")
	}
}

func TestCoinsUpdate(t *testing.T) {

	inputVal := 15000
	reward := 500

	input := models.NewUser(1193132712, inputVal, 0)
	log.Println("Saldo ingresado: ", inputVal)
	_, err = database.CreateNewUser(ctx, db, input)
	failOnError(t, err)
	log.Println("Se ha subido al servidor un usuario con saldo : ", inputVal)
	inputVal = inputVal + reward
	log.Println("Saldo modificado: ", inputVal)
	input.SetUserCoins(inputVal)

	_, err = database.UsersUpdateQuery(ctx, db, input)
	failOnError(t, err)
	log.Println("Se ha actualizado el saldo.")

	output, err := database.GetUserByID(ctx, db, 1193132712)
	failOnError(t, err)

	outputVal := output.GetUserCoins()
	log.Println("Saldo servidor: ", outputVal)

	if !reflect.DeepEqual(outputVal, inputVal) {
		t.Error("Input difers from output")
	}
}

func TestCoinsAccum(t *testing.T) {

	inputVal := 15000
	reward := 500

	inputUsr := models.NewUser(1193132715, inputVal, 0)
	inputAtt := models.NewAttraction("Bingo", "Espectacular Bingo mágico.", 150, 30, 0, 1)

	_, err = database.CreateNewUser(ctx, db, inputUsr)
	failOnError(t, err)
	log.Println("Se ha subido al servidor un usuario con saldo : ", inputVal)

	_, err = database.CreateNewAttraction(ctx, db, inputAtt)
	failOnError(t, err)

	outputAtt, err := database.GetAttractionByName(ctx, db, "Bingo")
	failOnError(t, err)

	log.Println("Gracias por montar la atracción", outputAtt.GetAttractionName())
	log.Println("Has recibido nuevos puntos: ", reward)

	inputVal = inputVal + reward
	log.Println("Saldo modificado: ", inputVal)
	inputUsr.SetUserCoins(inputVal)

	_, err = database.UsersUpdateQuery(ctx, db, inputUsr)
	failOnError(t, err)
	log.Println("Se ha actualizado el saldo.")

	outputUsr, err := database.GetUserByID(ctx, db, 1193132712)
	failOnError(t, err)

	outputVal := outputUsr.GetUserCoins()
	log.Println("Saldo servidor: ", outputVal)

	if !reflect.DeepEqual(outputVal, inputVal) {
		t.Error("Input difers from output")
	}
}

func TestAllAttractionView(t *testing.T) {

	input1 := models.NewAttraction("Ruleta de la suerte", "Es una gran ruleta", 150, 30, 0, 0)
	input2 := models.NewAttraction("Canal del amor", "Un romántico paseo en bote para los más tortolitos", 260, 20, 0, 1)

	_, err = database.CreateNewAttraction(ctx, db, input1)
	failOnError(t, err)
	_, err = database.CreateNewAttraction(ctx, db, input2)
	failOnError(t, err)

	output, err := database.GetAllAttractions(ctx, db)
	failOnError(t, err)
	log.Println(output)

}

func TestMeanTimeView(t *testing.T) {

	inputMeanTime := 120
	input := models.NewAttraction("Kamikaze", "Atracción que pondrá a prueba tus nervios.", inputMeanTime, 20, 0, 1)

	_, err = database.CreateNewAttraction(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetAttractionByName(ctx, db, "Kamikaze")
	failOnError(t, err)

	outputMeanTime := output.GetAttractionDuration()

	if !reflect.DeepEqual(outputMeanTime, inputMeanTime) {
		t.Error("Input difers from output")
	}

	log.Println("Tiempo de espera promedio:", outputMeanTime)

}

func TestTurnsView(t *testing.T) {

	input := models.NewAttraction("Casa de los espejos", "Piérdete con los espejos.", 150, 30, 24, 25)

	_, err = database.CreateNewAttraction(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetAttractionByName(ctx, db, "Casa de los espejos")
	failOnError(t, err)

	turns_available := output.GetAttractionCapacity() - output.GetAttractionCurrentTurn()

	log.Println(turns_available)

}

func TestGetTurn(t *testing.T) {

	inputAtt := models.NewAttraction("Carritos chocones", "Choca a todos los que puedas.", 150, 30, 12, 13)
	inputUsr := models.NewUser(1193132714, 16900, 0)

	_, err = database.CreateNewAttraction(ctx, db, inputAtt)
	failOnError(t, err)
	_, err = database.CreateNewUser(ctx, db, inputUsr)
	failOnError(t, err)

	outputAtt, err := database.GetAttractionByName(ctx, db, "Carritos chocones")
	failOnError(t, err)

	currentTurn := outputAtt.GetAttractionCurrentTurn()
	nextTurn := currentTurn + 1

	outputAtt.SetAttractionCurrentTurn(nextTurn)

	inputUsr.SetUserTurn(currentTurn)

	_, err = database.UsersUpdateQuery(ctx, db, inputUsr)
	failOnError(t, err)

	outputUsr, err := database.GetUserByID(ctx, db, 1193132714)
	failOnError(t, err)

	outputCurrentTurn := outputUsr.GetUserTurn()

	if !reflect.DeepEqual(outputUsr, inputUsr) {
		t.Error("Input difers from output")
	}
	if !reflect.DeepEqual(outputCurrentTurn, currentTurn) {
		t.Error("Input difers from output")
	}
	log.Println("turno tomado:", currentTurn)
	log.Println("actual turno del usuario:", outputCurrentTurn)
}
