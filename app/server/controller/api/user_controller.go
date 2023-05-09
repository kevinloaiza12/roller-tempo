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

func (uc *UserController) Users(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (uc *UserController) CreateNewUser(ctx echo.Context) error {
	type createUserRequest struct {
		Identification int `json:"identification"`
		Coins          int `json:"coins"`
		Turn           int `json:"turn"`
		Attraction     int `json:"attraction"`
	}

	var request createUserRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Identification <= 0 || request.Coins <= 0 || request.Turn <= 0 || request.Attraction <= 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	user := model.User{
		Identification: request.Identification,
		Coins:          request.Coins,
		Turn:           request.Turn,
		Attraction:     request.Attraction,
	}

	err = uc.userService.CreateUser(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "User created successfully"})
}

func (uc *UserController) GetUserByID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUserTurnAndAttraction(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	attractionID, err := strconv.Atoi(ctx.Param("attractionID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	err = uc.userService.UpdateUserTurnAndAttraction(id, attractionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "User turn and attraction updated successfully"})
}
