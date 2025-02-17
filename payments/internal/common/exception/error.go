package exception

import "github.com/gofiber/fiber/v2"

// define error here
var (
	ErrOrderNotFound    = fiber.NewError(fiber.StatusNotFound, "order is not found")
	ErrUserUnauthorized = fiber.NewError(fiber.StatusUnauthorized, "user unauthorized")
	ErrPaymentFailed    = fiber.NewError(fiber.StatusBadRequest, "payment failed")

	// generic error
	ErrInternalServerError = fiber.ErrInternalServerError
	ErrMessageNotPublished = fiber.NewError(fiber.StatusBadRequest, "message not published")
)
