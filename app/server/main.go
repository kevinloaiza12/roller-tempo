package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"roller-tempo/routes"
	"roller-tempo/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	envErr := godotenv.Load("config.env")
	if envErr != nil {
		log.Fatal(envErr)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: dbConn}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	utils.AutoMigrate(gormDb)

	app := echo.New()
	ctx := context.Background()
	routes.Register(app, ctx, gormDb)
	app.Logger.Fatal(app.Start(":3000"))
}
