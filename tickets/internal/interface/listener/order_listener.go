package listener

import (
	"context"
	"encoding/json"
	"log"
	"ticketing/tickets/internal/common/event"
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/model"
	"ticketing/tickets/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	QueueGroupName = "tickets-service"
)

type OrderListener struct {
	TicketUsecase usecase.TicketUsecase
	NatsConn      *nats.Conn
	Logger        *logrus.Logger
}

func NewOrderListener(ticketUseCase usecase.TicketUsecase, conn *nats.Conn, log *logrus.Logger) *OrderListener {
	return &OrderListener{
		TicketUsecase: ticketUseCase,
		NatsConn:      conn,
		Logger:        log,
	}
}

func (ol *OrderListener) HandleOrderCreated(data []byte) error {
	var event model.OrderCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal order created event")
		return err
	}

	log.Printf("Processing OrderCreatedEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateTicketRequest := new(model.UpdateTicketRequest)
	updateTicketRequest.Title = ""
	updateTicketRequest.Price = 0
	updateTicketRequest.OrderID = &event.ID
	updateTicketRequest.UserID = event.UserID
	updateTicketRequest.ID = event.TicketID

	if _, err := ol.TicketUsecase.Update(ctx, updateTicketRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to update ticket order")
		return err
	}

	return nil
}

func (ol *OrderListener) HandleOrderCancelled(data []byte) error {
	var event model.OrderCancelledEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal order created event")
		return err
	}

	log.Printf("Processing OrderCancelledEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateTicketRequest := new(model.UpdateTicketRequest)
	updateTicketRequest.Title = ""
	updateTicketRequest.Price = 0
	updateTicketRequest.ID = event.TicketID
	updateTicketRequest.UserID = event.UserID
	updateTicketRequest.OrderID = nil

	if _, err := ol.TicketUsecase.Update(ctx, updateTicketRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to update ticket order")
		return err
	}

	return nil
}

func (ol *OrderListener) Listen() {
	orderCreatedListener := &event.Listener{
		Subject:       domain.OrderCreated,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleOrderCreated,
	}

	orderCancelledListener := &event.Listener{
		Subject:       domain.OrderCancelled,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleOrderCancelled,
	}

	go orderCreatedListener.Listener()
	go orderCancelledListener.Listener()
}
