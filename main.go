package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON([]fiber.Map{
			{"id": 1, "name": "John Doe"},
			{"id": 2, "name": "Jane Doe"},
		})
	})

	app.Listen(":3000")
}
