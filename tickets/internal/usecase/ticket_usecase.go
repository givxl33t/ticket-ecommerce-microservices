package usecase

import (
	"context"
	"database/sql"
	"encoding/json"

	"ticketing/tickets/internal/common/exception"
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/infrastructure/event"
	"ticketing/tickets/internal/model"
	"ticketing/tickets/internal/model/mapper"
	"ticketing/tickets/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TicketUsecase interface {
	Create(ctx context.Context, request *model.CreateTicketRequest) (*model.TicketResponse, error)
	Update(ctx context.Context, request *model.UpdateTicketRequest) (*model.TicketResponse, error)
	FindAll(ctx context.Context) ([]model.TicketResponse, error)
	FindById(ctx context.Context, id int32) (*model.TicketResponse, error)
}

type TicketUsecaseImpl struct {
	TicketRepository repository.TicketRepository
	EventPublisher   *event.Publisher
	Logger           *logrus.Logger
	Validate         *validator.Validate
	Config           *viper.Viper
}

func NewTicketUsecase(ticketRepo repository.TicketRepository, log *logrus.Logger,
	validate *validator.Validate, config *viper.Viper) TicketUsecase {
	return &TicketUsecaseImpl{
		TicketRepository: ticketRepo,
		Logger:           log,
		Validate:         validate,
		Config:           config,
	}
}

func (uc *TicketUsecaseImpl) Create(ctx context.Context, request *model.CreateTicketRequest) (*model.TicketResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	ticket := new(domain.Ticket)
	ticket.Title = request.Title
	ticket.Price = request.Price
	ticket.UserID = request.UserID
	ticket.OrderID = sql.NullInt32{Valid: false}

	if err := uc.TicketRepository.Create(ctx, ticket); err != nil {
		uc.Logger.WithError(err).Error("failed create ticket to database")
		return nil, exception.ErrInternalServerError
	}

	// handling event logic
	event := model.TicketCreatedEvent{
		ID:      ticket.ID,
		Title:   ticket.Title,
		Price:   ticket.Price,
		UserID:  ticket.UserID,
		OrderID: nil,
	}

	data, err := json.Marshal(event)
	if err != nil {
		uc.Logger.WithError(err).Error("failed marshal event")
		return nil, exception.ErrInternalServerError
	}

	if err := uc.EventPublisher.Publish(domain.TicketCreated, data); err != nil {
		uc.Logger.WithError(err).Error("failed publish event TicketCreated event")
	}

	return mapper.ToTicketResponse(ticket), nil
}

func (uc *TicketUsecaseImpl) Update(ctx context.Context, request *model.UpdateTicketRequest) (*model.TicketResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	ticket, err := uc.TicketRepository.FindById(ctx, request.ID)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find ticket by id")
		return nil, exception.ErrTicketNotFound
	}

	if ticket.OrderID.Valid {
		uc.Logger.WithError(err).Error("failed find ticket by id")
		return nil, exception.ErrTicketAlreadyOrdered
	}

	if ticket.UserID != request.UserID {
		uc.Logger.WithError(err).Error("failed find ticket by id")
		return nil, exception.ErrUserUnauthorized
	}

	ticket.Title = request.Title
	ticket.Price = request.Price
	if request.OrderID == nil {
		ticket.OrderID = sql.NullInt32{Valid: false}
	} else {
		ticket.OrderID = sql.NullInt32{Int32: *request.OrderID, Valid: true}
	}

	if err := uc.TicketRepository.Update(ctx, ticket); err != nil {
		uc.Logger.WithError(err).Error("failed update ticket to database")
		return nil, exception.ErrInternalServerError
	}

	// handling event logic
	event := model.TicketUpdatedEvent{
		ID:      ticket.ID,
		Title:   ticket.Title,
		Price:   ticket.Price,
		UserID:  ticket.UserID,
		OrderID: ticket.OrderID,
	}

	data, err := json.Marshal(event)
	if err != nil {
		uc.Logger.WithError(err).Error("failed marshal event")
		return nil, exception.ErrInternalServerError
	}

	if err := uc.EventPublisher.Publish(domain.TicketUpdated, data); err != nil {
		uc.Logger.WithError(err).Error("failed publish event TicketUpdated event")
	}

	return mapper.ToTicketResponse(ticket), nil
}

func (uc *TicketUsecaseImpl) FindAll(ctx context.Context) ([]model.TicketResponse, error) {
	tickets, err := uc.TicketRepository.FindAll(ctx)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find all ticket to database")
		return nil, exception.ErrInternalServerError
	}

	responses := make([]model.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		responses[i] = *mapper.ToTicketResponse(&ticket)
	}

	return responses, nil
}

func (uc *TicketUsecaseImpl) FindById(ctx context.Context, id int32) (*model.TicketResponse, error) {
	ticket, err := uc.TicketRepository.FindById(ctx, id)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find ticket by id")
		return nil, exception.ErrTicketNotFound
	}

	return mapper.ToTicketResponse(ticket), nil
}
