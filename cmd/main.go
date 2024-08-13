package main

import (
	"ci/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	// Loging all incoming requests
	app.Use(logger.New())

	app.Get("/", handlers.GetHandler)

	app.Listen(":3000")

}
