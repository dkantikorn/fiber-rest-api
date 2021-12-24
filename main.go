package main

import (
	"log"

	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/dkantikorn/fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

//@Description Geting for welcome endpoint
//@Tags Welcome
//@Accept json
//@Product json
//@Success 200 {string} string
//@router /api [get]
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to awsome Go fiber rest api")
}

//Function making for setting internal routes
func SetupRoutes(app *fiber.App) {
	apiVersion := "/v1"
	apiURL := "/api"
	//Welcome endpoint
	app.Get(apiVersion+apiURL, welcome)

	//==============================================================================
	//User endpoints
	//==============================================================================
	user := app.Group(apiVersion + apiURL + "/users")
	user.Post("/", routes.CreateUser)
	user.Get("/", routes.GetUsers)
	user.Get("/:id", routes.GetUser)
	user.Put("/:id", routes.UpdateUser)
	user.Delete("/:id", routes.DeleteUser)

	//==============================================================================
	//Product enpoints
	//==============================================================================
	product := app.Group(apiVersion + apiURL + "/products")
	product.Post("/", routes.CreateProduct)
	product.Get("/", routes.GetProducts)
	product.Get("/:id", routes.GetProduct)
	product.Put("/:id", routes.UpdateProduct)
	product.Delete("/:id", routes.DeleteProduct)
}

func main() {
	//Connect to the daatbase
	database.ConnectDatabase()

	//Fiber new instance
	app := fiber.New()

	SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
