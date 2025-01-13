package handler

import (
	"strconv"
	"ticketing/orders/internal/model"
	"ticketing/orders/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	OrderUsecase usecase.OrderUsecase
	Logger       *logrus.Logger
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase, log *logrus.Logger) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: orderUsecase,
		Logger:       log,
	}
}

// Get Orders godoc
// @Summary Get All Orders
// @Description Fetches all order data.
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} []model.OrderResponse
// @Security Session
// @Router /orders [get]
func (h *OrderHandler) GetAll(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.AccessTokenPayload)

	response, err := h.OrderUsecase.FindAll(c.Context(), auth.ID)
	if err != nil {
		h.Logger.WithError(err).Error("error find all order")
		return err
	}
	return c.
		JSON(response)
}

// Create Order godoc
// @Summary Create Order
// @Description Create a new order.
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body model.CreateOrderRequest true "Order Create Request"
// @Success 201 {object} model.OrderResponse
// @Security Session
// @Router /orders [post]
func (h *OrderHandler) Create(c *fiber.Ctx) error {
	createOrderRequest := new(model.CreateOrderRequest)
	if err := c.BodyParser(createOrderRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	// do populate struct after validation or else it replaces the struct values
	auth := c.Locals("auth").(*model.AccessTokenPayload)
	createOrderRequest.UserID = auth.ID

	response, err := h.OrderUsecase.Create(c.Context(), createOrderRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error create order")
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(response)
}

// Get Order By ID godoc
// @Summary Get A Order
// @Description Fetches a order by id.
// @Tags Orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.OrderResponse
// @Security Session
// @Router /orders/{id} [get]
func (h *OrderHandler) GetByID(c *fiber.Ctx) error {
	authenticatedOrderRequest := new(model.AuthenticatedOrderRequest)

	// do populate struct after validation or else it replaces the struct values
	auth := c.Locals("auth").(*model.AccessTokenPayload)
	authenticatedOrderRequest.UserID = auth.ID

	// parse id as int32
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing id")
		return err
	}
	idint32 := int32(id)
	authenticatedOrderRequest.ID = idint32

	response, err := h.OrderUsecase.FindById(c.Context(), authenticatedOrderRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error find order by id")
		return err
	}

	return c.
		JSON(response)
}

// Cancel Order godoc
// @Summary Cancels A Order
// @Description  a order by id.
// @Tags Orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 204
// @Security Session
// @Router /orders/{id} [delete]
func (h *OrderHandler) Cancel(c *fiber.Ctx) error {
	authenticatedOrderRequest := new(model.AuthenticatedOrderRequest)

	// do populate struct after validation or else it replaces the struct values
	auth := c.Locals("auth").(*model.AccessTokenPayload)
	authenticatedOrderRequest.UserID = auth.ID

	// parse id as int32
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing id")
		return err
	}
	idint32 := int32(id)
	authenticatedOrderRequest.ID = idint32

	response, err := h.OrderUsecase.Cancel(c.Context(), authenticatedOrderRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error delete order by id")
		return err
	}

	return c.
		Status(fiber.StatusNoContent).
		JSON(response)
}
