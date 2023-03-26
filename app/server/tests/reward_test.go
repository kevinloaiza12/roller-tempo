package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func TestReward(t *testing.T) {

	ctx := context.Background()
	envErr := godotenv.Load("../config.env")
	if envErr != nil {
		os.Exit(1)
	}

	log.Println("Aquí cargó la config!")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		os.Exit(1)
	}

	log.Println("Aquí ya se conectó a la base de datos!")

	input := models.NewReward("Peluche", "Es un lindo peluche", 1235)

	_, err = database.CreateNewReward(ctx, db, input)
	failOnError(t, err)

	log.Println("Aquí ya hizo el primer query (crear)!")

	output, err := database.GetRewardByName(ctx, db, "Peluche")
	failOnError(t, err)

	log.Println("Aquí ya hizo el segundo query (obtener)!")

	if !reflect.DeepEqual(output, input) {
		t.Error("input difers from output")
	}
}

func TestPostReward(t *testing.T) {

	requestBody, _ := json.Marshal(map[string]interface{}{
		"id":          12345,
		"name":        "Peluchito",
		"description": "Es un buen peluche",
		"price":       15,
	})

	request, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/rewardregister", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("Error al enviar solicitud: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Código de estado esperado %d pero se recibió: %d", http.StatusOK, response.StatusCode)
	}

	var responseBody ResponseBody
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Error al decodificar respuesta: %v", err)
	}

	if responseBody.Message != controllers.OkMessageRegistry {
		t.Errorf("El valor de 'message' esperado era distinto, se recibió: %s", responseBody.Message)
	}
}
