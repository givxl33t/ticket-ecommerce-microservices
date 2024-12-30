package usecase

import (
	"context"
	"time"

	"ticketing/orders/internal/common/exception"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
	"ticketing/orders/internal/model/mapper"
	"ticketing/orders/internal/publisher"
	"ticketing/orders/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type OrderUsecase interface {
	Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error)
	FindAll(ctx context.Context, userId string) ([]model.OrderResponse, error)
	FindById(ctx context.Context, request *model.AuthenticatedOrderRequest) (*model.OrderResponse, error)
	Cancel(ctx context.Context, request *model.AuthenticatedOrderRequest) (*model.OrderResponse, error)
	UpdateStatus(ctx context.Context, request *model.UpdateOrderStatusRequest) error
}

type OrderUsecaseImpl struct {
	OrderRepository  repository.OrderRepository
	TicketRepository repository.TicketRepository
	OrderPublisher   publisher.OrderPublisher
	Logger           *logrus.Logger
	Validate         *validator.Validate
	Config           *viper.Viper
}

func NewOrderUsecase(orderRepo repository.OrderRepository, ticketRepo repository.TicketRepository, orderPublisher publisher.OrderPublisher, log *logrus.Logger,
	validate *validator.Validate, config *viper.Viper) OrderUsecase {
	return &OrderUsecaseImpl{
		OrderRepository:  orderRepo,
		TicketRepository: ticketRepo,
		OrderPublisher:   orderPublisher,
		Logger:           log,
		Validate:         validate,
		Config:           config,
	}
}

func (uc *OrderUsecaseImpl) Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	// Find the ticket the user id trying to order in the database
	ticket, err := uc.TicketRepository.FindById(ctx, request.TicketID)
	if err != nil {
		uc.Logger.WithError(err).Error("ticket not found")
		return nil, exception.ErrTicketNotFound
	}

	// check if the ticket is reserved in an order somewhere
	isReserved, err := uc.OrderRepository.IsTicketReserved(ctx, request.TicketID)
	if isReserved {
		uc.Logger.WithError(err).Error("ticket is reserved")
		return nil, exception.ErrTicketReserved
	}

	// calculate an expiration date for this order
	currentTime := time.Now()
	expiration := currentTime.Add(time.Second * time.Duration(domain.ExpirationWindowSeconds))

	// build the order and save it to the database
	order := new(domain.Order)
	order.UserID = request.UserID
	order.TicketID = request.TicketID
	order.Status = domain.Created
	order.ExpiresAt = expiration.Unix()

	if err := uc.OrderRepository.Create(ctx, order); err != nil {
		uc.Logger.WithError(err).Error("failed create order to database")
		return nil, exception.ErrInternalServerError
	}

	// publish an event to subject order:created
	if err := uc.OrderPublisher.Created(order, ticket); err != nil {
		uc.Logger.WithError(err).Error("failed publish event OrderCreated event")
		return nil, exception.ErrMessageNotPublished
	}

	return mapper.ToOrderResponse(order), nil
}

func (uc *OrderUsecaseImpl) FindAll(ctx context.Context, userId string) ([]model.OrderResponse, error) {
	orders, err := uc.OrderRepository.FindAll(ctx, userId)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find all ticket to database")
		return nil, exception.ErrInternalServerError
	}

	responses := make([]model.OrderResponse, len(orders))
	for i, order := range orders {
		responses[i] = *mapper.ToOrderResponse(&order)
	}

	return responses, nil
}

func (uc *OrderUsecaseImpl) FindById(ctx context.Context, request *model.AuthenticatedOrderRequest) (*model.OrderResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	order, err := uc.OrderRepository.FindById(ctx, request.ID)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find order by id")
		return nil, exception.ErrOrderNotFound
	}

	if order.UserID != request.UserID {
		uc.Logger.WithError(err).Error("user unauthorized to fetch")
		return nil, exception.ErrUserUnauthorized
	}

	return mapper.ToOrderResponse(order), nil
}

func (uc *OrderUsecaseImpl) Cancel(ctx context.Context, request *model.AuthenticatedOrderRequest) (*model.OrderResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	order, err := uc.OrderRepository.FindById(ctx, request.ID)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find order by id")
		return nil, exception.ErrOrderNotFound
	}

	if order.UserID != request.UserID {
		uc.Logger.WithError(err).Error("user unauthorized to delete")
		return nil, exception.ErrUserUnauthorized
	}

	// update the order to have status of cancelled
	order.Status = domain.Cancelled
	if err := uc.OrderRepository.Update(ctx, order); err != nil {
		uc.Logger.WithError(err).Error("failed cancel order in database")
		return nil, exception.ErrInternalServerError
	}

	// publish an event to subject order:cancelled
	if err := uc.OrderPublisher.Cancelled(order); err != nil {
		uc.Logger.WithError(err).Error("failed publish event OrderCancelled event")
		return nil, exception.ErrMessageNotPublished
	}

	return mapper.ToOrderResponse(order), nil
}

func (uc *OrderUsecaseImpl) UpdateStatus(ctx context.Context, request *model.UpdateOrderStatusRequest) error {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return err
	}

	order, err := uc.OrderRepository.FindById(ctx, request.ID)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find order by id")
		return exception.ErrOrderNotFound
	}

	order.Status = request.Status
	if err := uc.OrderRepository.Update(ctx, order); err != nil {
		uc.Logger.WithError(err).Error("failed update order in database")
		return exception.ErrInternalServerError
	}

	if request.Status == domain.Cancelled {
		// publish an event to subject order:cancelled
		if err := uc.OrderPublisher.Cancelled(order); err != nil {
			uc.Logger.WithError(err).Error("failed publish event OrderCancelled event")
			return exception.ErrMessageNotPublished
		}
	}

	return nil
}
