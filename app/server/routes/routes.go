package routes

import (
	"context"
	controller "roller-tempo/controller/api"

	"github.com/labstack/echo/v4"
)

func RegisterAttractionRoutes(app *echo.Echo, ctx context.Context, controller *controller.AttractionController) {
	app.GET("/api/attractions", controller.Attractions)
	app.GET("/api/attractions/:id", controller.GetAttractionByID)
	app.GET("/api/attractions/:id/turns", controller.GetNextRoundTurns)

	app.POST("/api/attractions/register", controller.CreateNewAttraction)
}

func RegisterRewardRoutes(app *echo.Echo, ctx context.Context, controller *controller.RewardController) {
	app.GET("/api/rewards", controller.Rewards)
	app.GET("/api/rewards/:id", controller.GetRewardByID)

	app.POST("/api/rewards/register", controller.CreateNewReward)
}

func RegisterUserRoutes(app *echo.Echo, ctx context.Context, controller *controller.UserController) {
	app.GET("/api/users", controller.Users)
	app.GET("/api/users/:id", controller.GetUserByID)

	app.POST("/api/users/register", controller.CreateNewUser)
	app.POST("/api/users/turn", controller.UpdateUserTurn)
	app.POST("/api/users/:id/reward", controller.RewardUser)
	app.POST("/api/users/:id/penalize", controller.PenalizeUser)

	app.PUT("/api/users/:id/removeturn", controller.RemoveUserTurn)
}
