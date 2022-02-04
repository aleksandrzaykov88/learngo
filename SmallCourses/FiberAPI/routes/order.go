package routes

import (
	"errors"
	"fiber_api/database"
	"fiber_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)
	responseOrders := []Order{}
	for _, order := range orders {
		var user models.User
		if err := findUser(order.UserRefer, &user); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		var product models.Product
		if err := findProduct(order.ProductRefer, &product); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateOrder struct {
		User    User    `json:"user"`
		Product Product `json:"product"`
	}

	var updateData UpdateOrder

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	order.User = user
	order.Product = product

	database.Database.Db.Save(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&order).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted Order")
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}
