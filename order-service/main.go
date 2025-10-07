package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/synapsis/common/config"
	"time"
)

func main() {
	app := fiber.New()
	cfg := config.FromEnv()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok")
	})

	app.Post("/orders", func(c *fiber.Ctx) error {
		var body struct {
			CustomerID string `json:"customer_id"`
			Items      []struct {
				SKU string `json:"sku"`
				Qty int    `json:"qty"`
			} `json:"items"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "Order received successfully",
			"customer_id": body.CustomerID,
			"order_id":    fmt.Sprintf("order-%d", time.Now().Unix()),
			"status":      "CONFIRMED",
		})
	})

	app.Listen(cfg.HTTPAddress)
}
