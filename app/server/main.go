package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(cors.New())
	routes.Register(app, ctx, db)

	app.Listen(":3000")
	fmt.Println("Server listening on port 3000")
}
