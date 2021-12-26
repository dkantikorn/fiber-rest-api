package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/dkantikorn/fiber-rest-api/config"
	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/dkantikorn/fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/dkantikorn/fiber-rest-api/docs"
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
	//not required AuthorizationRequired
	app.Get(apiVersion+apiURL, welcome)

	//Public resource in static route
	app.Static("/", "./public")

	//Swagger endpoint making for API document
	app.Get("/v1/api/swagger/*", swagger.Handler)

	//==============================================================================
	//User endpoints
	//==============================================================================
	user := app.Group(apiVersion + apiURL + "/users")

	user.Post("/login", routes.Auth)
	user.Post("/", routes.CreateUser)

	//AuthorizationRequired Action
	app.Use(routes.AuthorizationRequired())

	//need AuthorizationRequired
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

	//==============================================================================
	//Order enpoints
	//==============================================================================
	orderRoute := app.Group(apiVersion + apiURL + "/orders")
	orderRoute.Post("/", routes.CreateOrder)
	orderRoute.Get("/", routes.GetOrders)
	orderRoute.Get("/:id", routes.GetOrder)
	//end AuthorizationRequired
}

func main() {
	//Connect to the daatbase
	database.ConnectDatabase()
	config.InitJWT() //Init

	appPort := config.Config("APP_PORT")
	//Fiber new instance
	app := fiber.New()

	//Used logger for keeppink request logging ID
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	SetupRoutes(app)
	log.Fatal(app.Listen(":" + appPort))
}
