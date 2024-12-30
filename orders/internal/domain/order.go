package domain

const (
	Created         = "created"
	Cancelled       = "cancelled"
	AwaitingPayment = "awaiting:payment"
	Complete        = "complete"
)

const ExpirationWindowSeconds = 1 * 60

type Order struct {
	ID        int32  `gorm:"primaryKey;autoIncrement;column:id"`
	Status    string `gorm:"column:status;type:enum('created','cancelled','awaiting:payment','complete');not null"`
	UserID    string `gorm:"column:user_id"`
	TicketID  int32  `gorm:"column:ticket_id"`
	ExpiresAt int64  `gorm:"column:expires_at"`
	CreatedAt int64  `gorm:"column:created_at;autoUpdateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime:milli"`
}
