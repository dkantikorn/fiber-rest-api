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
