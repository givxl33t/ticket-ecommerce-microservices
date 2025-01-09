package mapper

import (
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
)

func ToOrderResponse(order *domain.Order) *model.OrderResponse {
	return &model.OrderResponse{
		ID:     order.ID,
		Status: order.Status,
		UserID: order.UserID,
		Ticket: model.TicketResponse{
			ID:    order.Ticket.ID,
			Title: order.Ticket.Title,
			Price: order.Ticket.Price,
		},
		ExpiresAt: order.ExpiresAt,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
