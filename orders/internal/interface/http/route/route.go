package route

import (
	"ticketing/orders/internal/interface/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, orderHandler *handler.OrderHandler, authMiddleware fiber.Handler) {
	prefixRouter := app.Group("/api")
	prefixRouter.Get("/orders", authMiddleware, orderHandler.GetAll)
	prefixRouter.Post("/orders", authMiddleware, orderHandler.Create)
	prefixRouter.Get("/orders/:id", authMiddleware, orderHandler.GetByID)
	prefixRouter.Delete("/orders/:id", authMiddleware, orderHandler.Cancel)
}
