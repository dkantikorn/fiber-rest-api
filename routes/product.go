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

//Function making for create a product infomation to the database
//@Summary Insert product info into the database
//@Description Create for product
//@Tags product
//@Accept json
//@Product json
//@Param name body string true "ProductName"
//@Param serial_number body string true "ProductSerialNumber"
//@Success 200 {object} entities.Product "product data with success message"
//@Failure 500 {object} entities.Product "product null data with failure message"
//@Router /v1/api/products [post]
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	database.Database.DB.Create(&product)
	responseProduct := utilities.CreateResponseProduct(product)
	return c.Status(utilities.CODE_SUCCESS).JSON(utilities.ReturnResponseMessageProduct("OK", "Product has been create successfully.", responseProduct))
}

//Function making for get all product information
//@Summary Get all product information
//@Description Get all product information
//@Tags product
//@Accept json
//@Product json
//@Success 200 {object} []entities.Product "array of product with success message"
//@Failure 201 {object} entities.Product "null product entities with error message"
//@Router /v1/api/products [get]
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.DB.Find(&products)
	if len(products) < 1 {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageProduct("FAILED", errors.New("product not present!!!").Error(), entities.Product{}))
	}

	var responseProducts []entities.Product
	for _, product := range products {
		tmpResponseProduct := utilities.CreateResponseProduct(product)
		responseProducts = append(responseProducts, tmpResponseProduct)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseProducts})
}

//Function maling fine for product info
func FindProduct(id int, product *models.Product) error {
	if id == 0 || id < 1 {
		return errors.New("Invalid for id: parameter!!!")
	}

	database.Database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product dose not exist")
	}

	return nil
}

//Function making for get product infomation
//@Summary Get for product information
//@Description Get for product information
//@Tags product
//@Accept json
//@Product json
//@Param id path integer true "Product ID"
//@Success 200 {object} entities.Product "product entities with success message"
//@Failure 201 {object} entities.Product "product entities with null value and error message"
//@Router /v1/api/products/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	var product models.Product
	if err := FindProduct(id, &product); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	responseProduct := utilities.CreateResponseProduct(product)
	return c.Status(http.StatusOK).JSON(utilities.ReturnResponseMessageProduct("OK", "OK", responseProduct))
}

//Function making for update product information
//@Summary Update for product infomation
//@Decription Product update infomation
//@Tags product
//@Accept json
//@Product json
//@Param id path integer true "ProductID"
//@Param name body string true "ProductName"
//@Param serial_number body string true "ProductSerialNumber"
//@Success 200 {object} entities.Product "product entities with success message"
//@Failure 201 {object} entities.Product "product entities with error message"
//@Router /v1/api/products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	var product models.Product
	if err := FindProduct(id, &product); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateProductData UpdateProduct
	if err := c.BodyParser(&updateProductData); err != nil {
		c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	product.Name = updateProductData.Name
	product.SerialNumber = updateProductData.SerialNumber

	database.Database.DB.Save(&product)
	responseProduct := utilities.CreateResponseProduct(product)
	return c.Status(http.StatusOK).JSON(utilities.ReturnResponseMessageProduct("OK", "OK", responseProduct))
}

//Function maling for delete product
//@Summary Delete for the product
//@Decription Delete for the product
//@Tags product
//@Accept json
//@Product json
//@Param id path integer true "ProductID"
//@Success 200 {object} entities.Product "product entities with success message"
//@Failure 201 {object} entities.Product "product entities with error message"
//@Router /v1/api/products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	var product models.Product
	if err := FindProduct(id, &product); err != nil {
		return c.Status(http.StatusNoContent).JSON(utilities.ReturnResponseMessageProduct("FAILED", err.Error(), entities.Product{}))
	}

	database.Database.DB.Delete(&product)
	return c.Status(http.StatusOK).JSON(utilities.ReturnResponseMessageProduct("OK", "Product has been deleted.", entities.Product{}))
}
