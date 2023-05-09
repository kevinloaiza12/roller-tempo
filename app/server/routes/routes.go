package routes

import (
	"context"
	controller "roller-tempo/controller/api"

	"github.com/labstack/echo/v4"
)

func RegisterAttractionRoutes(app *echo.Echo, ctx context.Context, controller *controller.AttractionController) {
	app.GET("/attractions", controller.Attractions)
}

func RegisterRewardRoutes(app *echo.Echo, ctx context.Context, controller *controller.RewardController) {
	app.GET("/rewards", controller.Rewards)
}

func RegisterUserRoutes(app *echo.Echo, ctx context.Context, controller *controller.UserController) {
	app.GET("/users", controller.Users)
}

/*

	MISSING ROUTES:

	app.GET("/api/userinfo/:id", controllers.GetUserInfo(ctx, db))
	app.GET("/api/attractioninfo/:name", controllers.GetAttractionInfo(ctx, db))
	app.GET("/api/rewardinfo/:name", controllers.GetRewardInfo(ctx, db))

	app.POST("/api/userregister/", controllers.PostUserRegister(ctx, db))
	app.POST("/api/usernextturn/", controllers.PostUserNextTurn(ctx, db))
	app.POST("/api/attractionregister/", controllers.PostAttractionRegister(ctx, db))
	app.POST("/api/rewardregister/", controllers.PostRewardRegister(ctx, db))
*/
