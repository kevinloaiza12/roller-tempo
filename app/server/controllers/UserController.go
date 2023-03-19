package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}
