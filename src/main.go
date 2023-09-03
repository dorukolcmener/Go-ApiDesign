package main

import (
	"apidesign/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Mount("/api", routes.Api())
	app.Listen(":3000")
}
