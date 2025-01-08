package unit

import (
	"context"
	"testing"

	"ticketing/payments/config"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/model"
	"ticketing/payments/internal/usecase"
	"ticketing/payments/test/unit/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	ctx = context.Background()
)

func TestCreatePayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	orderRepository := mocks.NewMockOrderRepository(ctrl)
	paymentRepository := mocks.NewMockPaymentRepository(ctrl)
	paymentPublisher := mocks.NewMockPaymentPublisher(ctrl)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository, paymentPublisher, orderRepository, logrus.New(), validator.New(), config.New())

	t.Run("success", func(t *testing.T) {
		orderRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)
		paymentRepository.EXPECT().FindById(ctx, gomock.Any()).Return(&domain.Ticket{
			ID: 1, // Manually set the return value for mocks??
		}, nil)
		orderRepository.EXPECT().IsTicketReserved(ctx, gomock.Any()).Return(false, nil)
		paymentPublisher.EXPECT().Created(gomock.Any()).Return(nil)

		request := &model.CreateOrderRequest{
			TicketID: 1,
			UserID:   "user-1",
		}

		response, err := paymentUsecase.Create(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, request.TicketID, response.Ticket.ID)
		assert.Equal(t, request.UserID, response.UserID)
	})

	t.Run("failed validation", func(t *testing.T) {
		request := &model.CreateOrderRequest{
			TicketID: 0,
			UserID:   "",
		}

		_, err := paymentUsecase.Create(ctx, request)
		assert.Error(t, err)
	})
}
