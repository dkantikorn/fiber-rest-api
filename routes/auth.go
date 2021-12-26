package routes

import (
	"time"

	"github.com/dkantikorn/fiber-rest-api/config"
	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/dkantikorn/fiber-rest-api/entities"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/utils"
	"github.com/golang-jwt/jwt/v4"
)

type (
	MsgLogin entities.Login
	MsgToken entities.Token
)

//Function making for autorize the username / password
//@Description User authorization
//@Tags user
//@Accept application/json
//@Product application/json
//@Param username body string true "Username"
//@Param password body string true "Password"
//@Success 200 {array} entities.Token
//@Failure 401 {string} string
//@router /v1/api/users/login [post]
func Auth(c *fiber.Ctx) error {
	var login MsgLogin
	err := c.BodyParser(&login)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
			"msg":   err.Error(),
		})
		return nil
	}
	// if login.Username != "sarawutt.b" || login.Password != "password" {
	// 	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Bad Credentials",
	// 	})
	// 	return nil
	// }

	if _, err := FindAuthLogin(entities.Login{Username: login.Username, Password: login.Password}); err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return nil
	}

	var user entities.User
	//database.Database.DB.Select([]string{"first_name", "last_name"}).Find(&user, "username = ?", login.Username)
	database.Database.DB.Find(&user, "username = ?", login.Username)

	token, err := createToken(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "StatusInternalServerError",
			"msg":   err.Error(),
		})
		return nil
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"user":          login.Username,
	})
	return nil
}

func createToken(userModel entities.User) (MsgToken, error) {
	var msgToken MsgToken
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = utils.UUID()
	claims["name"] = userModel.FirstName + " " + userModel.LastName
	// claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expired after 1 hour
	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.AccessToken = t
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = utils.UUID()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.RefreshToken = rt
	return msgToken, nil
}

//https://github.com/gofiber/jwt
func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// Filter:         nil,
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(config.JWTSecret),
		// SigningKeys:   nil,
		SigningMethod: "HS256",
		// ContextKey:    nil,
		// Claims:        nil,
		// TokenLookup:   nil,
		// AuthScheme:    nil,
	})
}

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sub": sub,
	})
	return nil
}
