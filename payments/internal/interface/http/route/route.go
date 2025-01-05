package route

import (
	"ticketing/payments/internal/interface/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, paymentHandler *handler.PaymentHandler, authMiddleware fiber.Handler) {
	prefixRouter := app.Group("/api")
	prefixRouter.Post("/payments", authMiddleware, paymentHandler.Create)
}
