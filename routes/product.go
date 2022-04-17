package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rixtrayker/fiber1/controllers"
)

func productRoutes(app *fiber.App) {

	// Product endpoints
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
}
