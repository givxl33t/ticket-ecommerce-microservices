package repository

import (
	"context"

	"ticketing/orders/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	Update(ctx context.Context, order *domain.Order) error
	FindAll(ctx context.Context, userId string) ([]domain.Order, error)
	FindById(ctx context.Context, id int32) (*domain.Order, error)
	IsTicketReserved(ctx context.Context, ticketId int32) (bool, error)
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

func (r *OrderRepositoryImpl) FindAll(ctx context.Context, userId string) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) FindById(ctx context.Context, id int32) (*domain.Order, error) {
	order := new(domain.Order)

	if err := r.DB.WithContext(ctx).Where("id = ?", id).Take(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepositoryImpl) Delete(ctx context.Context, id int32) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&domain.Order{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *OrderRepositoryImpl) IsTicketReserved(ctx context.Context, ticketId int32) (bool, error) {
	var order domain.Order
	if err := r.DB.WithContext(ctx).
		Where("ticket_id = ? AND status != ?", ticketId, domain.Cancelled).
		Take(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
