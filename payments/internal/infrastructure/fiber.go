package infrastructure

import (
	"fmt"
	"time"

	_ "ticketing/payments/docs"
	"ticketing/payments/internal/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: middleware.NewErrorHandler(),
		Prefork:      config.GetBool("APP_PREFORK"),
		WriteTimeout: config.GetDuration("APP_TIMEOUT") * time.Second,
		ReadTimeout:  config.GetDuration("APP_TIMEOUT") * time.Second,
	})

	app.Get("/api/payments/swagger/*", swagger.HandlerDefault)
	fmt.Println("Swagger UI available at /api/payments/swagger/")

	app.Use(recover.New())

	return app
}
