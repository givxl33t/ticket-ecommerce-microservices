package model

type CreateOrderRequest struct {
	ID     int32  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	Price  int64  `json:"price" validate:"required"`
}

type UpdateOrderStatusRequest struct {
	ID     int32  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type OrderCreatedEvent struct {
	ID       int32  `json:"id"`
	Status   string `json:"status"`
	UserID   string `json:"user_id"`
	TicketID int32  `json:"ticket_id"`
	Ticket   struct {
		ID    int32  `json:"id"`
		Title string `json:"title"`
		Price int64  `json:"price"`
	} `json:"ticket"`
	ExpiresAt int64 `json:"expires_at"`
}

type OrderCancelledEvent struct {
	ID       int32  `json:"id"`
	TicketID int32  `json:"ticket_id"`
	UserID   string `json:"user_id"`
}

type PaymentCreatedEvent struct {
	ID       int32  `json:"id"`
	OrderID  int32  `json:"order_id"`
	StripeID string `json:"stripe_id"`
}
