package domain

type Payment struct {
	ID       int32  `gorm:"primaryKey;autoIncrement;column:id"`
	OrderID  int32  `gorm:"column:order_id"`
	StripeID string `gorm:"column:stripe_id"`
}
