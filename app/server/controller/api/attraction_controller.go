package controller

import (
	"net/http"
	"roller-tempo/model"
	"roller-tempo/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AttractionController struct {
	attractionService *service.AttractionService
}

func NewAttractionController(attractionService *service.AttractionService) *AttractionController {
	return &AttractionController{attractionService: attractionService}
}

func (ac *AttractionController) Attractions(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (ac *AttractionController) CreateNewAttraction(ctx echo.Context) error {
	type createAttractionRequest struct {
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Duration         int     `json:"duration"`
		Capacity         int     `json:"capacity"`
		CurrentRoundTurn int     `json:"currentRoundTurn"`
		NextTurn         int     `json:"nextTurn"`
		PosX             float64 `json:"x"`
		PosY             float64 `json:"y"`
	}

	var request createAttractionRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Name == "" || request.Description == "" || request.Duration <= 0 ||
		request.Capacity <= 0 || request.CurrentRoundTurn < 0 || request.NextTurn < 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	attraction := model.Attraction{
		Name:             request.Name,
		Description:      request.Description,
		Duration:         request.Duration,
		Capacity:         request.Capacity,
		CurrentRoundTurn: request.CurrentRoundTurn,
		NextTurn:         request.NextTurn,
		PosX:             request.PosX,
		PosY:             request.PosY,
	}

	err = ac.attractionService.CreateAttraction(&attraction)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Attraction created successfully"})
}

func (ac *AttractionController) GetAttractionByID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	attraction, err := ac.attractionService.GetAttractionByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, attraction)
}

func (ac *AttractionController) GetNextRoundTurns(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	turns, err := ac.attractionService.GetNextRoundTurns(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, turns)
}
