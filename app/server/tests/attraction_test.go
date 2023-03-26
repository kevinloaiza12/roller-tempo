package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func TestAttraction(t *testing.T) {

	input := models.NewAttraction("Ruleta de la suerte", "Es una gran ruleta", 150, 30, 0)

	_, err = database.CreateNewAttraction(ctx, db, input)
	failOnError(t, err)

	output, err := database.GetAttractionByName(ctx, db, "Ruleta de la suerte")
	failOnError(t, err)

	if !reflect.DeepEqual(output, input) {
		t.Error("Input difers from output")
	}
}

func TestPostAttractionRegister(t *testing.T) {

	requestBody, _ := json.Marshal(map[string]interface{}{
		"name":        "Disney",
		"description": "Juego de Disney",
		"duration":    15,
		"capacity":    25,
		"nextTurn":    1,
	})

	request, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/attractionregister", bytes.NewBuffer(requestBody))
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
	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Error al decodificar respuesta: %v", err)
	}

	if responseBody.Message != controllers.OkMessageRegistry {
		t.Errorf("El valor de 'message' esperado era distinto, se recibió: %s", responseBody.Message)
	}
}
