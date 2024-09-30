package controllers

import (
	"github.com/Johnman67112/backend-dex/internal/core"
	"github.com/Johnman67112/backend-dex/internal/infra"
	"github.com/gofiber/fiber/v2"
)

func RegisterPokemon(c *fiber.Ctx) {
	var pokemon core.Pokemon

	if err := c.BodyParser(pokemon); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})

		return
	}

	if err := core.PokemonValidator(&pokemon); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validator error",
		})
	}

	infra.DB.Create(&pokemon)

	c.JSON(pokemon)
}

func GetPokemon(c *fiber.Ctx) {
	var Pokemon []core.Pokemon
	infra.DB.Find(&Pokemon)
	c.JSON(Pokemon)
}

func GetOnePokemon(c *fiber.Ctx) {
	var pokemon core.Pokemon
	number := c.Get("number")

	infra.DB.First(&pokemon, number)

	if pokemon.Number == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "no pokemon",
		})
	}

	c.JSON(pokemon)
}

func DeletePokemon(c *fiber.Ctx) {
	var pokemon core.Pokemon
	number := c.Get("number")

	infra.DB.Delete(&pokemon, number)

	c.JSON("Pokemon deleted sucessfully")
}

func EditPokemon(c *fiber.Ctx) {
	var pokemon core.Pokemon
	number := c.Get("number")

	infra.DB.First(&pokemon, number)

	if err := c.BodyParser(pokemon); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})

		return
	}

	if err := core.PokemonValidator(&pokemon); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

		return
	}

	infra.DB.Model(&pokemon).UpdateColumns(pokemon)
	c.JSON(pokemon)
}
