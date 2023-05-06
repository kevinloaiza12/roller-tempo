package routes

import (
	"context"
	controller "roller-tempo/controller/api"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(app *echo.Echo, ctx context.Context, db *gorm.DB) {
	app.GET("/attractions", controller.Attractions)
	app.GET("/rewards", controller.Rewards)
	app.GET("/users", controller.Users)

	/*
		app.GET("/api/userinfo/:id", controllers.GetUserInfo(ctx, db))
		app.GET("/api/attractioninfo/:name", controllers.GetAttractionInfo(ctx, db))
		app.GET("/api/rewardinfo/:name", controllers.GetRewardInfo(ctx, db))

		app.POST("/api/userregister/", controllers.PostUserRegister(ctx, db))
		app.POST("/api/usernextturn/", controllers.PostUserNextTurn(ctx, db))
		app.POST("/api/attractionregister/", controllers.PostAttractionRegister(ctx, db))
		app.POST("/api/rewardregister/", controllers.PostRewardRegister(ctx, db))
	*/
}
