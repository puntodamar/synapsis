package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/synapsis/common/config"
)

func main() {
	app := fiber.New()
	cfg := config.FromEnv()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok")
	})

	app.Listen(cfg.HTTPAddress)
}
