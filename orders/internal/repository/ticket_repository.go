package repository

import (
	"context"

	"ticketing/orders/internal/domain"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ctx context.Context, order *domain.Ticket) error
	Update(ctx context.Context, order *domain.Ticket) error
	FindById(ctx context.Context, id int32) (*domain.Ticket, error)
}

type TicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &TicketRepositoryImpl{DB: db}
}

func (r *TicketRepositoryImpl) Create(ctx context.Context, order *domain.Ticket) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TicketRepositoryImpl) Update(ctx context.Context, order *domain.Ticket) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TicketRepositoryImpl) FindById(ctx context.Context, id int32) (*domain.Ticket, error) {
	order := new(domain.Ticket)

	if err := r.DB.WithContext(ctx).Where("id = ?", id).Take(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
