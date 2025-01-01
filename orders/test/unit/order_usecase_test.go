package unit

import (
	"context"
	"testing"

	"ticketing/orders/config"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
	"ticketing/orders/internal/usecase"
	"ticketing/orders/test/unit/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	ctx = context.Background()
)

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	orderRepository := mocks.NewMockOrderRepository(ctrl)
	ticketRepository := mocks.NewMockTicketRepository(ctrl)
	orderPublisher := mocks.NewMockOrderPublisher(ctrl)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, ticketRepository, orderPublisher, logrus.New(), validator.New(), config.New())

	t.Run("success", func(t *testing.T) {
		orderRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)
		ticketRepository.EXPECT().FindById(ctx, gomock.Any()).Return(&domain.Ticket{
			ID: 1, // Manually set the return value for mocks??
		}, nil)
		orderRepository.EXPECT().IsTicketReserved(ctx, gomock.Any()).Return(false, nil)
		orderPublisher.EXPECT().Created(gomock.Any()).Return(nil)

		request := &model.CreateOrderRequest{
			TicketID: 1,
			UserID:   "user-1",
		}

		response, err := orderUsecase.Create(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, request.TicketID, response.Ticket.ID)
		assert.Equal(t, request.UserID, response.UserID)
	})

	t.Run("failed validation", func(t *testing.T) {
		request := &model.CreateOrderRequest{
			TicketID: 0,
			UserID:   "",
		}

		_, err := orderUsecase.Create(ctx, request)
		assert.Error(t, err)
	})
}
