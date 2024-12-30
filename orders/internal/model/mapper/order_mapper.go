package mapper

import (
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
)

func ToOrderResponse(order *domain.Order) *model.OrderResponse {
	return &model.OrderResponse{
		ID:        order.ID,
		Status:    order.Status,
		UserID:    order.UserID,
		TicketID:  order.TicketID,
		ExpiresAt: order.ExpiresAt,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
