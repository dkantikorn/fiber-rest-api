package routes

import (
	"errors"

	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/dkantikorn/fiber-rest-api/entities"
	"github.com/dkantikorn/fiber-rest-api/models"
	"github.com/dkantikorn/fiber-rest-api/utilities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func FindAuthLogin(login entities.Login) (bool, error) {
	var user models.User
	if err := database.Database.DB.Debug().Model(&entities.User{}).Where("username = ?", login.Username).Take(&user).Error; err != nil {
		return false, err
	}
	if err := models.VerifyPassword(user.Password, login.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, errors.New("Password missmatch")
	}
	return true, nil
}

//@Description Create new infomation of user
//@Tags user
//@Accept application/json
//@Product application/json
//@Success 200 {array} models.User
//@router /v1/api/users [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}

	database.Database.DB.Create(&user)
	responseUser := utilities.CreateResponseUser(user)
	return c.Status(utilities.CODE_SUCCESS).JSON(utilities.ReturnResponseMessageUser("OK", "Create user successfully", responseUser))
}

//@Description Get all existing of users
//@Tags user
//@Accept application/json
//@Product application/json
//@Success 200 {array} models.User
//@router /v1/api/users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.DB.Find(&users)
	if len(users) < 1 {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", errors.New("user rows not found!!!").Error(), entities.User{}))
	}
	responseUsers := []entities.User{}
	for _, user := range users {
		tmpResponseUser := utilities.CreateResponseUser(user)
		responseUsers = append(responseUsers, tmpResponseUser)
	}
	return c.Status(utilities.CODE_SUCCESS).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseUsers})
}

//Function making for find the user with parameter is us id
func FindUser(id int, user *models.User) error {
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user dose not exist")
	}
	return nil
}

//@Description Get for user details
//@Tags user
//@Accept application/json
//@Product application/json
//@Success 200 {object} entities.User
//@router /v1/api/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", "Please ensure that id: parameter is integer", entities.User{}))
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(utilities.CODE_NOT_FOUND).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}
	responseUser := utilities.CreateResponseUser(user)
	return c.Status(utilities.CODE_SUCCESS).JSON(utilities.ReturnResponseMessageUser("OK", "OK", responseUser))
}

//Function making for update the user infomation
//@Description Update user information
//@Tags user
//@Accept application/json
//@Product application/json
//@Param first_name body string true "FirstName"
//@Param last_name body string true "LastName"
//@Success 200 {object} entities.User
//@router /v1/api/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	//Checking for id as parameter is valid
	if err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", "Please ensure that id: parameter is integer", entities.User{}))
	}
	//Check for existing user on the database
	if err := FindUser(id, &user); err != nil {
		return c.Status(utilities.CODE_NOT_FOUND).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}

	//Define data type for update user info
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateUserData UpdateUser
	if err := c.BodyParser(&updateUserData); err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}

	user.FirstName = updateUserData.FirstName
	user.LastName = updateUserData.LastName

	database.Database.DB.Save(&user)
	responseUser := utilities.CreateResponseUser(user)
	return c.Status(utilities.CODE_SUCCESS).JSON(utilities.ReturnResponseMessageUser("OK", "Update user successfully", responseUser))
}

//Function making for delete user information
//@Description Delete for user infomation
//@Tags user
//@Accept application/json
//@Product application/json
//@Success 200 {json} response
//@router /v1/api/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	//Checking for id is a lalid parameter
	if err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", "Please ensure that id: parameter is integer", entities.User{}))
	}
	//Checking for user infomation is existing
	if err := FindUser(id, &user); err != nil {
		return c.Status(utilities.CODE_NOT_FOUND).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}
	//Delete and tracking for error
	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return c.Status(utilities.CODE_FAILED).JSON(utilities.ReturnResponseMessageUser("FAILED", err.Error(), entities.User{}))
	}
	// return c.Status(200).JSON(fiber.Map{"status": "OK", "message": "User deleted Successfully", "data": nil})
	return c.Status(utilities.CODE_SUCCESS).JSON(utilities.ReturnResponseMessageUser("OK", "User deleted Successfully", entities.User{}))
}
