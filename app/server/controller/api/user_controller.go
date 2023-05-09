package controller

import (
	"net/http"
	mapper "roller-tempo/dto/mapper"
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
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, mapper.ToUserDTO(user))
}

func (uc *UserController) UpdateUserTurn(ctx echo.Context) error {
	type updateUserTurnRequest struct {
		UserID       int `json:"userID"`
		AttractionID int `json:"attractionID"`
	}

	var request updateUserTurnRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.UserID <= 0 || request.AttractionID <= 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	err = uc.userService.UpdateUserTurnAndAttraction(request.UserID, request.AttractionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "User turn updated successfully"})
}
