package mapper

import (
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/model"
)

func ToTicketResponse(ticket *domain.Ticket) *model.TicketResponse {
	var orderID *string

	// check whether the sql.NullString is valid
	if ticket.OrderID.Valid {
		orderID = &ticket.OrderID.String
	} else {
		orderID = nil
	}

	return &model.TicketResponse{
		ID:        ticket.ID,
		Title:     ticket.Title,
		Price:     ticket.Price,
		UserID:    ticket.UserID,
		OrderID:   orderID,
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
	}
}
