package controllers

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func Attractions(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}

func GetAttractionInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessage400))
		}
		result, err := database.GetAttractionByID(ctx, db, id)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
		return c.JSON(result.AttractionToJSON())
	}
}

func PostAttractionRegister(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type AttractionRegisterRequest struct {
			Id          int64  `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    int    `json:"duration"`
			Capacity    int    `json:"capacity"`
			NextTurn    int    `json:"nextTurn"`
		}

		var info AttractionRegisterRequest
		if err := c.BodyParser(&info); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
		}

		if _, attractionExists := database.GetAttractionByID(ctx, db, int(info.Id)); attractionExists != sql.ErrNoRows {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

		attraction := models.NewAttraction(info.Id, info.Name, info.Description, info.Duration, info.Capacity, info.NextTurn)
		if _, err := database.CreateNewAttraction(ctx, db, attraction); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, err.Error()))
		}

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}
