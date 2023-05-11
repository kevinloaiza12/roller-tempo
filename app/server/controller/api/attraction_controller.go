package controller

import (
	"net/http"
	controller "roller-tempo/controller/request"
	utils "roller-tempo/controller/utils"
	mapper "roller-tempo/dto/mapper"
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
	attractions, err := ac.attractionService.GetAllAttractions()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": attractions})
}

func (ac *AttractionController) CreateNewAttraction(ctx echo.Context) error {

	var request controller.CreateAttractionRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Name == "" || request.Description == "" || request.Duration <= 0 ||
		request.Capacity <= 0 || request.CurrentRoundTurn < 0 || request.NextTurn < 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": utils.BadRequest})
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
		ImagePath:        request.Name + ".jpg",
	}

	err = ac.attractionService.CreateAttraction(&attraction)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
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

	return ctx.JSON(http.StatusOK, mapper.ToAttractionDTO(attraction))
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
