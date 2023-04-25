package controllers

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/database"
	"github.com/kevinloaiza12/roller-tempo/app/models"
)

func Users(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}

func GetUserInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessage400))
		}
		result, err := database.GetUserByID(ctx, db, id)
		if err != nil {
			return c.JSON(fiber.NewError(fiber.StatusNotFound, ErrorMessage404))
		}
		return c.JSON(result.UserToJSON())
	}
}

func PostUserRegister(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type UserRegisterRequest struct {
			Id    int `json:"id"`
			Coins int `json:"coins"`
			Turn  int `json:"turn"`
            Attraction string `json:"attraction"`

		}

		var info UserRegisterRequest
		if err := c.BodyParser(&info); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
		}

		if _, userExists := database.GetUserByID(ctx, db, info.Id); userExists != sql.ErrNoRows {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

		user := models.NewUser(info.Id, info.Coins, info.Turn, info.Attraction)
		if _, err := database.CreateNewUser(ctx, db, user); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, err.Error()))
		}

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}

func PostUserNextTurn(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type UserTurnUpdate struct {
			Id    int `json:"id"`
			//Coins int `json:"coins"`
			//Turn  int `json:"turn"`
            Attraction string `json:"attraction"`

		}

		var info UserTurnUpdate
		if err := c.BodyParser(&info); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
		}
        name := info.Attraction

		user, userExists := database.GetUserByID(ctx, db, info.Id)
		//if user, userExists := database.GetUserByID(ctx, db, info.Id); userExists != nil {
        if userExists != nil {
			return c.JSON(fiber.NewError(fiber.StatusBadRequest, ErrorMessageRegisteredUser))
		}

        turn, errTurn := database.GetNextTurn(ctx, db, name)    
        if errTurn != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, ErrorMessage500))
        }

        user.SetUserTurn(turn)
        user.SetUserAttraction(name)

        if _, err := database.UsersUpdateQuery(ctx, db, user); err != nil {
			return c.JSON(fiber.NewError(fiber.StatusServiceUnavailable, ErrorMessage500))
        }

		return c.JSON(fiber.Map{
			"message": OkMessageRegistry,
		})
	}
}
