package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rixtrayker/fiber1/controllers"
)

func orderRoutes(app *fiber.App) {

	// Order endpoints
	app.Post("/api/orders", controllers.CreateOrder)
	app.Get("/api/orders", controllers.GetOrders)
	app.Get("/api/orders/:id", controllers.GetOrder)
}
