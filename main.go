package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", home)

	app.Listen(":3000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
