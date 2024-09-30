package ui

import (
	"github.com/Johnman67112/backend-dex/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func StartAPI() {
	app := fiber.New()

	app.Get("/pokemon", func(c *fiber.Ctx) error {
		controllers.GetPokemon(c)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/pokemon/:id", func(c *fiber.Ctx) error {
		controllers.GetOnePokemon(c)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/pokemon", func(c *fiber.Ctx) error {
		controllers.RegisterPokemon(c)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/pokemon/:id", func(c *fiber.Ctx) error {
		controllers.EditPokemon(c)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete("/pokemon/:id", func(c *fiber.Ctx) error {
		controllers.DeletePokemon(c)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Listen(":3000")
}
