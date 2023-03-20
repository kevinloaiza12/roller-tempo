package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func TestUser(t *testing.T) {

	input := models.NewUser(1193132710, 15000, 0)

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

func TestPostUser(t *testing.T) {

	requestBody, _ := json.Marshal(map[string]interface{}{
		"id":    100022,
		"coins": 4220,
		"turn":  12,
	})

	request, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/userregister", bytes.NewBuffer(requestBody))
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
