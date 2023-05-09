package controller

import (
	"net/http"
	"roller-tempo/service"

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
