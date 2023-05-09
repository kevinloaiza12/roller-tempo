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

func (ac *AttractionController) CreateNewAttraction(c echo.Context) error {
	type CreateAttractionRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Duration    int     `json:"duration"`
		Capacity    int     `json:"capacity"`
		CurrentTurn int     `json:"currentTurn"`
		NextTurn    int     `json:"nextTurn"`
		PosX        float64 `json:"x"`
		PosY        float64 `json:"y"`
	}

	var request CreateAttractionRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Name == "" || request.Description == "" || request.Duration <= 0 ||
		request.Capacity <= 0 || request.CurrentTurn < 0 || request.NextTurn < 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	attraction := model.Attraction{
		Name:        request.Name,
		Description: request.Description,
		Duration:    request.Duration,
		Capacity:    request.Capacity,
		CurrentTurn: request.CurrentTurn,
		NextTurn:    request.NextTurn,
		PosX:        request.PosX,
		PosY:        request.PosY,
	}

	err = ac.attractionService.CreateAttraction(&attraction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Attraction created successfully"})
}

func (ac *AttractionController) GetAttractionByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	attraction, err := ac.attractionService.GetAttractionByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, attraction)
}
