package handlers

import "github.com/gofiber/fiber/v2"

func GetHandler(c *fiber.Ctx) error {
	return c.SendString("Hello from server")
}
