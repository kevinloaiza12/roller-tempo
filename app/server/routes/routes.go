package routes

import (
	"context"
	controller "roller-tempo/controller/api"

	"github.com/labstack/echo/v4"
)

func RegisterAttractionRoutes(app *echo.Echo, ctx context.Context, controller *controller.AttractionController) {
	app.GET("/api/attractions", controller.Attractions)
	app.POST("/api/attractions/register", controller.CreateNewAttraction)
	app.GET("/api/attractions/:id", controller.GetAttractionByID)
	app.GET("/api/attractions/:id/turns", controller.GetNextRoundTurns)

}

func RegisterRewardRoutes(app *echo.Echo, ctx context.Context, controller *controller.RewardController) {
	app.GET("/api/rewards", controller.Rewards)
	app.POST("/api/rewards/register", controller.CreateNewReward)
	app.GET("/api/rewards/:id", controller.GetRewardByID)

}

func RegisterUserRoutes(app *echo.Echo, ctx context.Context, controller *controller.UserController) {
	app.GET("/api/users", controller.Users)
	app.POST("/api/users/register", controller.CreateNewUser)
	app.POST("/api/users/turn", controller.UpdateUserTurn)

	app.POST("/api/users/:id/reward", controller.RewardUser)
	app.POST("/api/users/:id/penalize", controller.PenalizeUser)
	app.POST("api/users/buyreward", controller.BuyReward)
	app.GET("/api/users/:id", controller.GetUserByID)

	app.PUT("/api/users/:id/removeturn", controller.RemoveUserTurn)
}
