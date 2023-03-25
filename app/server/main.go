package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"

	//"google.golang.org/genproto/googleapis/cloud/functions/v1"
	"github.com/joho/godotenv"
	"github.com/kevinloaiza12/roller-tempo/app/routes"
)

func main() {
	envErr := godotenv.Load("config.env")
	if envErr != nil {
		log.Fatal(envErr)
	}

	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
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
