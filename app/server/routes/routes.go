package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
)

func Register(app *fiber.App, ctx context.Context, db *sql.DB) {
	app.Get("/users", controllers.Users)
	app.Get("/attractions", controllers.GetAllAttractionsInfo(ctx, db))
	app.Get("/rewards", controllers.GetAllRewards(ctx, db))
	app.Get("/attractionturns/:name", controllers.GetAttractionTurns(ctx, db))
	app.Get("/api/userinfo/:id", controllers.GetUserInfo(ctx, db))
	app.Get("/api/attractioninfo/:name", controllers.GetAttractionInfo(ctx, db))
	app.Get("/api/rewardinfo/:name", controllers.GetRewardInfo(ctx, db))
	app.Get("/api/nextturn/:name", controllers.GetNextTurn(ctx, db))

	app.Post("/api/userregister/", controllers.PostUserRegister(ctx, db))
	app.Post("/api/attractionregister/", controllers.PostAttractionRegister(ctx, db))
	app.Post("/api/rewardregister/", controllers.PostRewardRegister(ctx, db))
}
