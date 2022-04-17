package main

import (
	"log"

	"github.com/rixtrayker/fiber1/database"
	"github.com/rixtrayker/fiber1/routes"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Get("/api", welcome)
	routes.OrderRoutes(app)
	routes.UserRoutes(app)
	routes.ProductRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
