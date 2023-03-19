package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	//"google.golang.org/genproto/googleapis/cloud/functions/v1"
	"github.com/kevinloaiza12/roller-tempo/app/routes"
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
	routes.Register(app)

	app.Listen(":3000")
	fmt.Println("Server listening on port 3000")
}
