package controller

import (
	"net/http"
	"roller-tempo/service"

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
