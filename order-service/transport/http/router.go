package http

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, h *Handler) {
	app.Post("/orders", h.CreateOrder)
}
