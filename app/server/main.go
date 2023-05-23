package main

import (
	"context"
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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	ctx := context.Background()

	envErr := godotenv.Load("config.env")
	if envErr != nil {
		log.Fatal(envErr)
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DBHost"),
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"),
		os.Getenv("DBName"),
		os.Getenv("DBPort"),
	)

	gormDb, err := gorm.Open(postgres.New(postgres.Config{DSN: connStr}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	utils.AutoMigrate(gormDb)

	attractionRepository := repository.NewAttractionRepository(gormDb)
	attractionService := service.NewAttractionService(attractionRepository)

	rewardRepository := repository.NewRewardRepository(gormDb)
	rewardService := service.NewRewardService(rewardRepository)

	userRepository := repository.NewUserRepository(gormDb)
	userService := service.NewUserService(userRepository, attractionService, rewardService)

	attractionController := controller.NewAttractionController(attractionService)
	rewardController := controller.NewRewardController(rewardService)
	userController := controller.NewUserController(userService)

	app := echo.New()
	routes.RegisterAttractionRoutes(app, ctx, attractionController)
	routes.RegisterRewardRoutes(app, ctx, rewardController)
	routes.RegisterUserRoutes(app, ctx, userController)
	app.Logger.Fatal(app.Start(":3000"))
}
