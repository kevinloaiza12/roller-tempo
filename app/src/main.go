package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	//"google.golang.org/genproto/googleapis/cloud/functions/v1"
)

const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBPassword = "secret"
	DBName     = "rollertempo"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "hola front, saludos desde el back",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server listening on port 3000")
}
