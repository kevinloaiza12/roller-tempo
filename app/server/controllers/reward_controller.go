package controllers

import (
	"context"
	"database/sql"
	"strconv"

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
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessage400))
		}
		result, err := database.GetRewardByID(ctx, db, id)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
		return c.JSON(result.RewardToJSON())
	}
}

func PostRewardRegister(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type RewardRegisterRequest struct {
			Id          int64  `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Price       int    `json:"price"`
		}

		var info RewardRegisterRequest
		if err := c.BodyParser(&info); err != nil {
			return err
		}

		if _, rewardExists := database.GetRewardByID(ctx, db, int(info.Id)); rewardExists != sql.ErrNoRows {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

		reward := models.NewReward(info.Id, info.Name, info.Description, info.Price)
		if _, err := database.CreateNewReward(ctx, db, reward); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, err.Error()))
		}

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}
