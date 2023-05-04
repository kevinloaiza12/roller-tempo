package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func Rewards(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}

func GetRewardInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		result, err := database.GetRewardByName(ctx, db, name)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
		return c.JSON(result.RewardToJSON())
	}
}

func GetAllRewards(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := database.GetAllRewards(ctx, db)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}

		return c.JSON(result)
	}
}

func PostRewardRegister(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type RewardRegisterRequest struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Price       int    `json:"price"`
		}

		var info RewardRegisterRequest
		if err := c.BodyParser(&info); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
		}

		if _, rewardExists := database.GetRewardByName(ctx, db, info.Name); rewardExists != sql.ErrNoRows {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

		reward := models.NewReward(info.Name, info.Description, info.Price)
		if _, err := database.CreateNewReward(ctx, db, reward); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, err.Error()))
		}

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}
