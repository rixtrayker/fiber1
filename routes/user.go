package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rixtrayker/fiber1/controllers"
)

func userRoutes(app *fiber.App) {

	// User endpoints
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}
