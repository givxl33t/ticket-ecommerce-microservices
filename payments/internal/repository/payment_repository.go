package repository

import (
	"context"

	"ticketing/payments/internal/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(ctx context.Context, order *domain.Payment) error
}

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{DB: db}
}

func (r *PaymentRepositoryImpl) Create(ctx context.Context, order *domain.Payment) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}
