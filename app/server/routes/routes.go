package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
)

func Register(app *fiber.App, ctx context.Context, db *sql.DB) {
	app.Get("/users", controllers.Users)
	app.Get("/api/userinfo/:id", controllers.UserInfo(ctx, db))
	app.Get("/api/attractioninfo/:id", controllers.AttractionInfo(ctx, db))
	app.Get("/api/rewardinfo/:id", controllers.RewardInfo(ctx, db))
}
