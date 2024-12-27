package model

type CreateTicketRequest struct {
	Title  string `json:"title" validate:"required,max=100"`
	Price  int64  `json:"price" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type UpdateTicketRequest struct {
	ID     int32  `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required,max=100"`
	Price  int64  `json:"price" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type UpdateTicketOrderRequest struct {
	ID      int32  `json:"id" validate:"required"`
	OrderID *int32 `json:"order_id"`
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
