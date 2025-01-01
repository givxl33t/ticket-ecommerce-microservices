package usecase

import (
	"context"
	"ticketing/orders/internal/common/exception"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
	"ticketing/orders/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TicketUsecase interface {
	Create(ctx context.Context, request *model.TicketRequest) error
	Update(ctx context.Context, request *model.TicketRequest) error
}

type TicketUsecaseImpl struct {
	TicketRepository repository.TicketRepository
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

func (uc *TicketUsecaseImpl) Create(ctx context.Context, request *model.TicketRequest) error {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return err
	}

	ticket := new(domain.Ticket)
	ticket.ID = request.ID
	ticket.Title = request.Title
	ticket.Price = request.Price

	if err := uc.TicketRepository.Create(ctx, ticket); err != nil {
		uc.Logger.WithError(err).Error("failed create ticket to database")
		return exception.ErrInternalServerError
	}

	return nil
}

func (uc *TicketUsecaseImpl) Update(ctx context.Context, request *model.TicketRequest) error {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return err
	}

	ticket, err := uc.TicketRepository.FindById(ctx, request.ID)
	if err != nil {
		uc.Logger.WithError(err).Error("ticket not found")
		return exception.ErrTicketNotFound
	}

	ticket.Title = request.Title
	ticket.Price = request.Price

	if err := uc.TicketRepository.Update(ctx, ticket); err != nil {
		uc.Logger.WithError(err).Error("failed update ticket to database")
		return exception.ErrInternalServerError
	}

	return nil
}
