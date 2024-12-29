package exception

import "github.com/gofiber/fiber/v2"

// define error here
var (
	ErrTicketNotFound       = fiber.NewError(fiber.StatusNotFound, "ticket is not found")
	ErrTicketAlreadyOrdered = fiber.NewError(fiber.StatusBadRequest, "ticket is already ordered")
	ErrUserUnauthorized     = fiber.NewError(fiber.StatusUnauthorized, "user unauthorized")

	// generic error
	ErrInternalServerError = fiber.ErrInternalServerError
	ErrMessageNotPublished = fiber.NewError(fiber.StatusBadRequest, "message not published")
)
