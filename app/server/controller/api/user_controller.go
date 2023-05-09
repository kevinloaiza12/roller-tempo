package controller

import (
	"net/http"
	"roller-tempo/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Users(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
