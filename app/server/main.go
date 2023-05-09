package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	controller "roller-tempo/controller/api"
	"roller-tempo/repository"
	"roller-tempo/routes"
	"roller-tempo/service"
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
	defer dbConn.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: dbConn}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	utils.AutoMigrate(gormDb)

	attractionRepository := repository.NewAttractionRepository(gormDb)
	attractionService := service.NewAttractionService(attractionRepository)
	attractionController := controller.NewAttractionController(attractionService)

	rewardRepository := repository.NewRewardRepository(gormDb)
	rewardService := service.NewRewardService(rewardRepository)
	rewardController := controller.NewRewardController(rewardService)

	userRepository := repository.NewUserRepository(gormDb)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app := echo.New()
	ctx := context.Background()
	routes.RegisterAttractionRoutes(app, ctx, attractionController)
	routes.RegisterRewardRoutes(app, ctx, rewardController)
	routes.RegisterUserRoutes(app, ctx, userController)
	app.Logger.Fatal(app.Start(":3000"))
}
