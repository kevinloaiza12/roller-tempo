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
	input := models.NewAttraction("Ruleta de la suerte", "Es una gran ruleta", 150, 30, 0, 0)

	_, err = database.CreateNewAttraction(ctx, db, input)
	if err != nil {
		t.Fatalf(err.Error())
	}

	output, err := database.GetAttractionByName(ctx, db, "Ruleta de la suerte")
	if err != nil {
		t.Fatalf(err.Error())
	}

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
		"currentTurn": 1,
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
