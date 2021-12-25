package routes

import (
	"errors"
	"net/http"

	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/dkantikorn/fiber-rest-api/entities"
	"github.com/dkantikorn/fiber-rest-api/models"
	"github.com/dkantikorn/fiber-rest-api/utilities"
	"github.com/gofiber/fiber/v2"
)

//Function making for fine order with specific orderID
func FindOrder(id int, order *models.Order) error {
	database.Database.DB.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order dose not exist")
	}
	return nil
}

//Function making for create a order infomation to the database
//@Summary Insert order info into the database
//@Description Create for order
//@Tags order
//@Accept json
//@Product json
//@Param product_id body string true "productId"
//@Param user_id body string true "userId"
//@Success 200 {object} entities.Order "order data with success message"
//@Failure 500 {object} entities.Order "order null data with failure message"
//@Router /v1/api/orders [post]
func CreateOrder(c *fiber.Ctx) error {

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	//Check for user information
	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	//Checking for product
	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	database.Database.DB.Create(&order)
	responseUser := utilities.CreateResponseUser(user)
	responseProduct := utilities.CreateResponseProduct(product)
	responseOrder := utilities.CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(http.StatusOK).JSON(utilities.ReturnResponseMessageOrder("OK", "Order has been created", responseOrder))
}

//Function making for get all order information
//@Summary Get all order information
//@Description Get all order information
//@Tags order
//@Accept json
//@Product json
//@Success 200 {object} []entities.Order "array of order with success message"
//@Failure 201 {object} entities.Order "null order entities with error message"
//@Router /v1/api/orders [get]
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.DB.Find(&orders)

	responseOrders := []entities.Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		if err := FindUser(order.UserRefer, &user); err != nil {
			return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
		}

		if err := FindProduct(order.ProductRefer, &product); err != nil {
			return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
		}

		tmpResponseOrder := utilities.CreateResponseOrder(order, utilities.CreateResponseUser(user), utilities.CreateResponseProduct(product))
		responseOrders = append(responseOrders, tmpResponseOrder)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseOrders})
}

//Function making for get order infomation
//@Summary Get for order information
//@Description Get for order information
//@Tags order
//@Accept json
//@Product json
//@Param id path integer true "order ID"
//@Success 200 {object} entities.Order "order entities with success message"
//@Failure 201 {object} entities.Order "order entities with null value and error message"
//@Router /v1/api/orders/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageOrder("FAILED", "Please ensure that :id is an integer", entities.Order{}))
	}

	var order models.Order
	if err := FindOrder(id, &order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	var user models.User
	var product models.Product
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageOrder("FAILED", err.Error(), entities.Order{}))
	}

	responseOrder := utilities.CreateResponseOrder(order,
		utilities.CreateResponseUser(user),
		utilities.CreateResponseProduct(product))
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseOrder})
}
