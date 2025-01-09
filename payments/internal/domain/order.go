package domain

const (
	Created         = "created"
	Cancelled       = "cancelled"
	AwaitingPayment = "awaiting:payment"
	Complete        = "complete"
)

type Order struct {
	ID     int32  `gorm:"primaryKey;column:id"`
	Status string `gorm:"column:status;type:enum('created','cancelled','awaiting:payment','complete');not null"`
	UserID string `gorm:"column:user_id"`
	Price  int64  `gorm:"column:price"`
}
