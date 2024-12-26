package repository

import (
	"context"

	"ticketing/tickets/internal/domain"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *domain.Ticket) error
	Update(ctx context.Context, ticket *domain.Ticket) error
	FindAll(ctx context.Context) ([]domain.Ticket, error)
	FindById(ctx context.Context, id uint) (*domain.Ticket, error)
}

type TicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &TicketRepositoryImpl{DB: db}
}

func (r *TicketRepositoryImpl) Create(ctx context.Context, ticket *domain.Ticket) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(ticket).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TicketRepositoryImpl) Update(ctx context.Context, ticket *domain.Ticket) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(ticket).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TicketRepositoryImpl) FindAll(ctx context.Context) ([]domain.Ticket, error) {
	var tickets []domain.Ticket

	if err := r.DB.WithContext(ctx).Where("order_id is NULL").Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepositoryImpl) FindById(ctx context.Context, id uint) (*domain.Ticket, error) {
	ticket := new(domain.Ticket)

	if err := r.DB.WithContext(ctx).Where("id = ?", id).Take(ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}
