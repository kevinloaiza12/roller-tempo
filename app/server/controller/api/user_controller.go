package controller

import (
	"net/http"
	"roller-tempo/service"
	"strconv"

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

func (uc *UserController) GetUserByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, user)
}
