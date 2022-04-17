package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rixtrayker/fiber1/database"
	"github.com/rixtrayker/fiber1/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	// Payment endpoints
	app.Post("/api/payments", routes.CreatePayment)
	app.Get("/api/payments", routes.GetPayments)
	app.Get("/api/payments/:id", routes.GetPayment)
	app.Put("/api/payments/:id", routes.UpdatePayment)
	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
