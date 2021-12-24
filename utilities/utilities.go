package utilities

import (
	"github.com/dkantikorn/fiber-rest-api/entities"
	"github.com/dkantikorn/fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

//Function making for build create response message for user
func CreateResponseUser(userModel models.User) entities.User {
	return entities.User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

//Function making for build the response message infomation to caller RestAPI
func ReturnResponseMessageUser(status string, message string, data entities.User) map[string]interface{} {
	if data.ID == 0 {
		return fiber.Map{"status": status, "message": message, "data": nil}
	} else {
		return fiber.Map{"status": status, "message": message, "data": data}
	}
}

//Function making for build the response message for the product
func CreateResponseProduct(productModel models.Product) entities.Product {
	return entities.Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

//Function making for build the response message for the production information
func ReturnResponseMessageProduct(status string, message string, data entities.Product) map[string]interface{} {
	if data.ID == 0 {
		return fiber.Map{"status": status, "message": message, "data": nil}
	}
	return fiber.Map{"status": status, "message": message, "data": data}
}
