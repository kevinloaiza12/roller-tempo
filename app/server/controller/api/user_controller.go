package controller

import (
	"net/http"
	"roller-tempo/model"
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

func (uc *UserController) CreateNewUser(c echo.Context) error {
	type CreateUserRequest struct {
		Identification int    `json:"identification"`
		Coins          int    `json:"coins"`
		Turn           int    `json:"turn"`
		Attraction     string `json:"attraction"`
	}

	var request CreateUserRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Identification <= 0 || request.Coins <= 0 || request.Turn <= 0 || request.Attraction == "" {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	user := model.User{
		Identification: request.Identification,
		Coins:          request.Coins,
		Turn:           request.Turn,
		Attraction:     request.Attraction,
	}

	err = uc.userService.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User created successfully"})
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
