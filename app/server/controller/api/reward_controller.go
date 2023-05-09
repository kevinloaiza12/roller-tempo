package controller

import (
	"net/http"
	controller "roller-tempo/controller/request"
	mapper "roller-tempo/dto/mapper"
	"roller-tempo/model"
	"roller-tempo/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RewardController struct {
	rewardService *service.RewardService
}

func NewRewardController(rewardService *service.RewardService) *RewardController {
	return &RewardController{rewardService: rewardService}
}

func (rc *RewardController) Rewards(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (rc *RewardController) CreateNewReward(ctx echo.Context) error {

	var request controller.CreateRewardRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Name == "" || request.Description == "" || request.Price <= 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	reward := model.Reward{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	err = rc.rewardService.CreateReward(&reward)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Reward created successfully"})
}

func (rc *RewardController) GetRewardByID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	reward, err := rc.rewardService.GetRewardByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, mapper.ToRewardDTO(reward))
}
