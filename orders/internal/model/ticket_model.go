package model

import "database/sql"

type TicketRequest struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
	Price int64  `json:"price" `
}

type TicketResponse struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Price     int64  `json:"price"`
	UserID    string `json:"user_id"`
	OrderID   *int32 `json:"order_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type TicketUpdatedEvent struct {
	ID      int32         `json:"id"`
	Title   string        `json:"title"`
	Price   int64         `json:"price"`
	UserID  string        `json:"user_id"`
	OrderID sql.NullInt32 `json:"order_id"`
}

type TicketCreatedEvent struct {
	ID      int32         `json:"id"`
	Title   string        `json:"title"`
	Price   int64         `json:"price"`
	UserID  string        `json:"user_id"`
	OrderID sql.NullInt32 `json:"order_id"`
}
