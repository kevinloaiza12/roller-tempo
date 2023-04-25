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

func TestUser(t *testing.T) {
	input := models.NewUser(1193132710, 15000, 0, "")

	_, err = database.CreateNewUser(ctx, db, input)
	if err != nil {
		t.Fatalf(err.Error())
	}

	output, err := database.GetUserByID(ctx, db, 1193132710)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(output, input) {
		t.Error("Input difers from output")
	}
}

func TestPostUser(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"id":    100022,
		"coins": 4220,
		"turn":  12,
        "attraction": "cars",
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
	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Error al decodificar respuesta: %v", err)
	}

	if responseBody.Message != controllers.OkMessageRegistry {
		t.Errorf("El valor de 'message' esperado era distinto, se recibió: %s", responseBody.Message)
	}
}

func TestCoinsUpdate(t *testing.T) {
	inputVal := 15000
	reward := 500

    input := models.NewUser(1193132716, inputVal, 0, "")
	_, err = database.CreateNewUser(ctx, db, input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	inputVal = inputVal + reward
	input.SetUserCoins(inputVal)

	_, err = database.UsersUpdateQuery(ctx, db, input)
	if err != nil {
		t.Fatalf(err.Error())
	}

	output, err := database.GetUserByID(ctx, db, 1193132716)
	if err != nil {
		t.Fatalf(err.Error())
	}

	outputVal := output.GetUserCoins()

	if !reflect.DeepEqual(outputVal, inputVal) {
		t.Error("Input difers from output")
	}
}

func TestNextTurn(t *testing.T) {
    initialTurn := 0
    userID := 1004626964
    user := models.NewUser(userID, 5000, initialTurn, "")
	_, err = database.CreateNewUser(ctx, db, user)
	if err != nil {
		t.Fatalf(err.Error())
	}

    attraction := models.NewAttraction("carrusel","es un carrusel", 500, 50, 0, 50, 2.0, 8.0)
    _, attraction_err := database.CreateNewAttraction(ctx, db, attraction)
    if attraction_err != nil {
		t.Fatalf(attraction_err.Error())
    }

	requestBody, _ := json.Marshal(map[string]interface{}{
		"id":         userID,
        "attraction": attraction.GetAttractionName(),
	})

	request, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/usernextturn", bytes.NewBuffer(requestBody))
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
    updatedUser, newUserErr := database.GetUserByID(ctx, db, userID)
    if newUserErr != nil {
        t.Errorf("Error al obtener usuario actualizado: %s", responseBody.Message)
    }

    if initialTurn == updatedUser.GetUserTurn() {
		t.Error("Initial turn does not change")
    }
}
