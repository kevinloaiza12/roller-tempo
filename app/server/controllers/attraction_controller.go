package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func Attractions(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}

func GetAllAttractionsInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := database.GetAllAttractions(ctx, db)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}

		return c.JSON(result)
	}
}

func GetAttractionInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		result, err := database.GetAttractionByName(ctx, db, name)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
		return c.JSON(result.AttractionToJSON())
	}
}

func GetNextTurn(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		result, err := database.GetAttractionByName(ctx, db, name)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
    currentTurn := result.GetAttractionCurrentTurn()
    nextTurn := currentTurn+1

    result.SetAttractionCurrentTurn(nextTurn)
    
    _, updateErr := database.AttractionsUpdateQuery(ctx, db, result)
		if updateErr != nil {
			return c.JSON("Error updating next turn in database")
		}

    turn := map[string]int{
        "turn": nextTurn,
    }
    return c.JSON(turn)
	}
}

func PostAttractionRegister(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type AttractionRegisterRequest struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    int    `json:"duration"`
			Capacity    int    `json:"capacity"`
			CurrentTurn int    `json:"currentTurn"`
			NextTurn    int    `json:"nextTurn"`
		}

		var info AttractionRegisterRequest
		if err := c.BodyParser(&info); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
		}

		if _, attractionExists := database.GetAttractionByName(ctx, db, info.Name); attractionExists != sql.ErrNoRows {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

		attraction := models.NewAttraction(info.Name, info.Description, info.Duration, info.Capacity, info.CurrentTurn, info.NextTurn)
		if _, err := database.CreateNewAttraction(ctx, db, attraction); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, err.Error()))
		}

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}
