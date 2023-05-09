package controller

import (
	"net/http"
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

func (rc *RewardController) Rewards(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (rc *RewardController) CreateNewReward(c echo.Context) error {
	type CreateRewardRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}

	var request CreateRewardRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if request.Name == "" || request.Description == "" || request.Price <= 0 {
		errorMessage := "Invalid request body. Please provide all the required fields."
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errorMessage})
	}

	reward := model.Reward{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	err = rc.rewardService.CreateReward(&reward)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Reward created successfully"})
}

func (rc *RewardController) GetRewardByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	reward, err := rc.rewardService.GetRewardByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return c.JSON(http.StatusOK, reward)
}
