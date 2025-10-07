package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	recovermw "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/synapsis/common/config"
	"github.com/synapsis/order-service/domain"
	"github.com/synapsis/order-service/repo"
	"github.com/synapsis/order-service/transport/http"
	"github.com/synapsis/order-service/transport/http/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(recovermw.New(recovermw.Config{EnableStackTrace: true}))

	cfg := config.FromEnv()

	gdb, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok")
	})

	orderRepo := repo.NewOrderGormRepo(gdb)
	service := domain.NewService(orderRepo)

	h := http.NewHandler(service)
	http.Router(app, h)

	err = app.Listen(cfg.HTTPAddress)
	if err != nil {
		return
	}
}
