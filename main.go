package main

import (
	"test_assement/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	api.SetupRoutes(app)

	app.Listen(":8080")

}
