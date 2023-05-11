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

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Users(ctx echo.Context) error {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": users})
}

func (uc *UserController) CreateNewUser(ctx echo.Context) error {
	var request controller.CreateUserRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Identification <= 0 || request.Coins <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": utils.BadRequest})
	}

	user := model.User{
		Identification: request.Identification,
		Coins:          request.Coins,
	}

	err = uc.userService.CreateUser(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
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
	var request controller.UpdateUserTurnRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.UserID <= 0 || request.AttractionID <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": utils.BadRequest})
	}

	err = uc.userService.UpdateUserTurnAndAttraction(request.UserID, request.AttractionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
}

func (uc *UserController) RemoveUserTurn(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err = uc.userService.RemoveTurn(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
}

func (uc *UserController) RewardUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	amount, err := strconv.Atoi(ctx.QueryParam("amount"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err = uc.userService.RewardUser(id, amount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
}

func (uc *UserController) PenalizeUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	amount, err := strconv.Atoi(ctx.QueryParam("amount"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err = uc.userService.PenalizeUser(id, amount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": utils.OK})
}
