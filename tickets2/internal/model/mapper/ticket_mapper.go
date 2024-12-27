package mapper

import (
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/model"
)

func ToTicketResponse(ticket *domain.Ticket) *model.TicketResponse {
	var orderID *int32

	// check whether the sql.NullInt32 is valid
	if ticket.OrderID.Valid {
		orderID = &ticket.OrderID.Int32
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
