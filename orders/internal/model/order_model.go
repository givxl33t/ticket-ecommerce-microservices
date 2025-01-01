package model

import "ticketing/orders/internal/domain"

type CreateOrderRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	TicketID int32  `json:"ticket_id" validate:"required"`
}

type UpdateOrderStatusRequest struct {
	ID     int32  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type AuthenticatedOrderRequest struct {
	ID     int32  `json:"id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type OrderResponse struct {
	ID        int32         `json:"id"`
	Status    string        `json:"status"`
	UserID    string        `json:"user_id"`
	Ticket    domain.Ticket `json:"ticket"`
	ExpiresAt int64         `json:"expires_at"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
}

type OrderCreatedEvent struct {
	ID        int32  `json:"id"`
	Status    string `json:"status"`
	UserID    string `json:"user_id"`
	TicketID  int32  `json:"ticket_id"`
	ExpiresAt int64  `json:"expires_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type OrderCancelledEvent struct {
	ID       int32  `json:"id"`
	TicketID int32  `json:"ticket_id"`
	UserID   string `json:"user_id"`
}

type ExpirationCompleteEvent struct {
	OrderID int32 `json:"order_id"`
}

type PaymentCreatedEvent struct {
	ID                int32  `json:"id"`
	OrderID           int32  `json:"order_id"`
	CheckoutSessionID string `json:"checkout_session_id"`
}
