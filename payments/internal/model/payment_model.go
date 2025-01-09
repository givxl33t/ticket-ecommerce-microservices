package model

type PaymentRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	OrderID int32  `json:"order_id" validate:"required"`
}

type PaymentResponse struct {
	ID string `json:"id"`
}
