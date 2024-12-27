package handler

import (
	"strconv"
	"ticketing/tickets/internal/model"
	"ticketing/tickets/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TicketHandler struct {
	TicketUsecase usecase.TicketUsecase
	Logger        *logrus.Logger
}

func NewTicketHandler(ticketUseCase usecase.TicketUsecase, log *logrus.Logger) *TicketHandler {
	return &TicketHandler{
		TicketUsecase: ticketUseCase,
		Logger:        log,
	}
}

// Get Tickets godoc
// @Summary Get All Tickets
// @Description Fetches all ticket data.
// @Tags Tickets
// @Accept json
// @Produce json
// @Success 200 {object} []model.TicketResponse
// @Router /tickets [get]
func (h *TicketHandler) GetAll(c *fiber.Ctx) error {
	response, err := h.TicketUsecase.FindAll(c.Context())
	if err != nil {
		h.Logger.WithError(err).Error("error find all ticket")
		return err
	}

	return c.
		JSON(response)
}

// Create Ticket godoc
// @Summary Create Ticket
// @Description Create a new ticket.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param request body model.CreateTicketRequest true "Ticket Create Request"
// @Success 201 {object} model.TicketResponse
// @Router /tickets [post]
func (h *TicketHandler) Create(c *fiber.Ctx) error {
	// auth := c.Locals("auth").(*model.AccessTokenPayload)
	createTicketRequest := new(model.CreateTicketRequest)

	createTicketRequest.UserID = "anything"

	if err := c.BodyParser(createTicketRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	response, err := h.TicketUsecase.Create(c.Context(), createTicketRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error create ticket")
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(response)
}

// Get Ticket By ID godoc
// @Summary Get A Ticket
// @Description Fetches a ticket by id.
// @Tags Tickets
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} model.TicketResponse
// @Router /tickets/{id} [get]
func (h *TicketHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")

	// parse id as int32
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing id")
		return err
	}

	idint32 := int32(id)

	response, err := h.TicketUsecase.FindById(c.Context(), idint32)
	if err != nil {
		h.Logger.WithError(err).Error("error find ticket by id")
		return err
	}

	return c.
		JSON(response)
}

// Update Ticket godoc
// @Summary Update A Ticket
// @Description Updates a ticket data by id.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param request body model.UpdateTicketRequest true "Ticket Update Request"
// @Param id path int true "Ticket ID"
// @Success 200 {object} model.TicketResponse
// @Router /tickets/{id} [put]
func (h *TicketHandler) Update(c *fiber.Ctx) error {
	updateTicketRequest := new(model.UpdateTicketRequest)
	if err := c.BodyParser(updateTicketRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	idStr := c.Params("id")

	// parse id as int32
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing id")
		return err
	}

	idint32 := int32(id)
	updateTicketRequest.ID = idint32

	response, err := h.TicketUsecase.Update(c.Context(), updateTicketRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error update ticket")
		return err
	}

	return c.
		JSON(response)
}
