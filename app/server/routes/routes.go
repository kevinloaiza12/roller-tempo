package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/controllers"
)

func Register(app *fiber.App) {
	app.Get("/users", controllers.Users)
}
