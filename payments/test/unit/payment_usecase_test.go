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
	"github.com/stripe/stripe-go/v74"
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
	paymentGateway := mocks.NewMockPaymentGateway(ctrl)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository, paymentPublisher, paymentGateway, orderRepository, logrus.New(), validator.New(), config.New())

	t.Run("success", func(t *testing.T) {
		orderRepository.EXPECT().FindById(ctx, gomock.Any()).Return(&domain.Order{
			ID:     1,
			UserID: "user-1",
		}, nil)
		paymentRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)
		paymentPublisher.EXPECT().Created(gomock.Any()).Return(nil)
		paymentGateway.EXPECT().CreatePayment(gomock.Any()).Return(&stripe.PaymentIntent{
			ClientSecret: "client-secret",
		}, nil)

		request := &model.PaymentRequest{
			UserID:  "user-1",
			OrderID: 1,
		}

		_, err := paymentUsecase.Create(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("failed validation", func(t *testing.T) {
		request := &model.PaymentRequest{
			UserID:  "",
			OrderID: 0,
		}

		_, err := paymentUsecase.Create(ctx, request)
		assert.Error(t, err)
	})
}
