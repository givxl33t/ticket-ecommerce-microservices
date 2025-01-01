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
		Ticket:    order.Ticket,
		ExpiresAt: order.ExpiresAt,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
