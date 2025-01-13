package handler

import (
	"ticketing/payments/internal/model"
	"ticketing/payments/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	PaymentUsecase usecase.PaymentUsecase
	Logger         *logrus.Logger
}

func NewPaymentHandler(paymentUsecase usecase.PaymentUsecase, log *logrus.Logger) *PaymentHandler {
	return &PaymentHandler{
		PaymentUsecase: paymentUsecase,
		Logger:         log,
	}
}

// Create Payment godoc
// @Summary Create Payment
// @Description Create a new payment.
// @Tags payments
// @Accept json
// @Produce json
// @Param request body model.PaymentRequest true "Payment Create Request"
// @Success 201 {object} model.PaymentResponse
// @Security Session
// @Router /payments [post]
func (h *PaymentHandler) Create(c *fiber.Ctx) error {
	createPaymentRequest := new(model.PaymentRequest)
	if err := c.BodyParser(createPaymentRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	// do populate struct after validation or else it replaces the struct values
	auth := c.Locals("auth").(*model.AccessTokenPayload)
	createPaymentRequest.UserID = auth.ID

	response, err := h.PaymentUsecase.Create(c.Context(), createPaymentRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error create order")
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(response)
}
