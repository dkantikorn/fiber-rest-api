package main

import (
	"log"

	"github.com/dkantikorn/fiber-rest-api/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to awsome Go fiber rest api")
}

func main() {
	//Connect to the daatbase
	database.ConnectDatabase()

	//Fiber new instance
	app := fiber.New()

	app.Get("/api", welcome)
	log.Fatal(app.Listen(":3000"))
}
