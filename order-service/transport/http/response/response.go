package response

import "github.com/gofiber/fiber/v2"

type Envelope map[string]any

func Success(c *fiber.Ctx, status int, data Envelope) error {
	out := Envelope{
		"status": "success",
		"code":   status,
	}
	for k, v := range data {
		out[k] = v
	}
	return c.Status(status).JSON(out)
}
