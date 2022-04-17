package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rixtrayker/fiber1/database"
	"github.com/rixtrayker/fiber1/models"
)

type Order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

func CreateResponseOrder(order models.Order, user User) Order {
	return Order{ID: order.ID, User: user}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := findUser(order.UserId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseOrder := CreateResponseOrder(order, responseUser)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", order.UserId)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	database.Database.Db.First(&user, order.UserId)
	responseUser := CreateResponseUser(user)

	responseOrder := CreateResponseOrder(order, responseUser)

	return c.Status(200).JSON(responseOrder)

}
