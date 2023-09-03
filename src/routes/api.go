package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Api() *fiber.App {
	api := fiber.New()
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api main page.")
	})
	// Mount User Route
	api.Mount("/user", User())
	// Return the Whole Route
	return api
}
