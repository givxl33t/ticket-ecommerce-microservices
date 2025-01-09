package usecase

import (
	"context"

	"ticketing/payments/internal/common/exception"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/model"
	"ticketing/payments/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type OrderUsecase interface {
	Create(ctx context.Context, request *model.CreateOrderRequest) error
	UpdateStatus(ctx context.Context, request *model.UpdateOrderStatusRequest) error
}

type OrderUsecaseImpl struct {
	OrderRepository repository.OrderRepository
	Logger          *logrus.Logger
	Validate        *validator.Validate
	Config          *viper.Viper
}

func NewOrderUsecase(orderRepo repository.OrderRepository, log *logrus.Logger,
	validate *validator.Validate, config *viper.Viper) OrderUsecase {
	return &OrderUsecaseImpl{
		OrderRepository: orderRepo,
		Logger:          log,
		Validate:        validate,
		Config:          config,
	}
}

func (uc *OrderUsecaseImpl) Create(ctx context.Context, request *model.CreateOrderRequest) error {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return err
	}

	order := new(domain.Order)
	order.ID = request.ID
	order.UserID = request.UserID
	order.Status = domain.Created
	order.Price = request.Price

	if err := uc.OrderRepository.Create(ctx, order); err != nil {
		uc.Logger.WithError(err).Error("failed create order to database")
		return exception.ErrInternalServerError
	}

	return nil
}

func (uc *OrderUsecaseImpl) UpdateStatus(ctx context.Context, request *model.UpdateOrderStatusRequest) error {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed to validate  request body")
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

	return nil
}
