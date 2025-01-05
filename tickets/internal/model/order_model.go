package model

import "ticketing/tickets/internal/domain"

type OrderCreatedEvent struct {
	ID        int32         `json:"id"`
	Status    string        `json:"status"`
	UserID    string        `json:"user_id"`
	Ticket    domain.Ticket `json:"ticket"`
	ExpiresAt int64         `json:"expires_at"`
}

type OrderCancelledEvent struct {
	ID        int32  `json:"id"`
	Status    string `json:"status"`
	UserID    string `json:"user_id"`
	TicketID  int32  `json:"ticket_id"`
	ExpiresAt int64  `json:"expires_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
