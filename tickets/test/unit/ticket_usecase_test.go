package unit

import (
	"context"
	"testing"

	"ticketing/tickets/config"
	"ticketing/tickets/internal/model"
	"ticketing/tickets/internal/usecase"
	"ticketing/tickets/test/unit/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	ctx = context.Background()
)

func TestCreateTicket(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepository := mocks.NewMockTicketRepository(ctrl)
	ticketPublisher := mocks.NewMockTicketPublisher(ctrl)
	ticketUsecase := usecase.NewTicketUsecase(ticketRepository, ticketPublisher, logrus.New(), validator.New(), config.New())

	t.Run("success", func(t *testing.T) {
		ticketRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)
		ticketPublisher.EXPECT().Created(gomock.Any()).Return(nil)

		request := &model.CreateTicketRequest{
			Title:  "concert",
			Price:  30000,
			UserID: "user-1",
		}

		response, err := ticketUsecase.Create(ctx, request)
		assert.NoError(t, err)
		assert.Equal(t, request.Title, response.Title)
		assert.Equal(t, request.Price, response.Price)
	})

	t.Run("failed validation", func(t *testing.T) {
		request := &model.CreateTicketRequest{
			Title: "",
			Price: 0,
		}

		_, err := ticketUsecase.Create(ctx, request)
		assert.Error(t, err)
	})
}
