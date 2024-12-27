package model

type OrderCreatedEvent struct {
	ID        *int32 `json:"id"`
	Status    string `json:"status"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
	Ticket    struct {
		ID    int32 `json:"id"`
		Price int64 `json:"price"`
	} `json:"ticket"`
}

type OrderCancelledEvent struct {
	ID     *int32 `json:"id"`
	Ticket struct {
		ID int32 `json:"id"`
	} `json:"ticket"`
}
