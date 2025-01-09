package repository

import (
	"context"

	"ticketing/payments/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	Update(ctx context.Context, order *domain.Order) error
	FindById(ctx context.Context, id int32) (*domain.Order, error)
}

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (r *OrderRepositoryImpl) Create(ctx context.Context, order *domain.Order) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *OrderRepositoryImpl) Update(ctx context.Context, order *domain.Order) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *OrderRepositoryImpl) FindById(ctx context.Context, id int32) (*domain.Order, error) {
	order := new(domain.Order)
	if err := r.DB.WithContext(ctx).Where("id = ?", id).Take(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
