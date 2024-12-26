package route

import (
	"ticketing/tickets/internal/interface/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, ticketHandler *handler.TicketHandler, authMiddleware fiber.Handler) {
	prefixRouter := app.Group("/api")
	prefixRouter.Get("/tickets", ticketHandler.GetAll)
	prefixRouter.Post("/tickets", ticketHandler.Create)
	// temporary comment prefixRouter.Post("/tickets", authMiddleware, ticketHandler.Create)
	prefixRouter.Get("/tickets/:id", ticketHandler.GetByID)
	prefixRouter.Put("/tickets/:id", authMiddleware, ticketHandler.Update)
}
