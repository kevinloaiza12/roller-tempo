package controller

import (
	"net/http"
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
