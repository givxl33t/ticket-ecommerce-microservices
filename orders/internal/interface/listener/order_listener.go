package listener

import (
	"context"
	"encoding/json"
	"log"
	"ticketing/orders/internal/common/event"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"
	"ticketing/orders/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	QueueGroupName = "orders-service"
)

type OrderListener struct {
	OrderUsecase  usecase.OrderUsecase
	TicketUsecase usecase.TicketUsecase
	NatsConn      *nats.Conn
	Logger        *logrus.Logger
}

func NewOrderListener(orderUseCase usecase.OrderUsecase, ticketUseCase usecase.TicketUsecase, conn *nats.Conn, log *logrus.Logger) *OrderListener {
	return &OrderListener{
		OrderUsecase:  orderUseCase,
		TicketUsecase: ticketUseCase,
		NatsConn:      conn,
		Logger:        log,
	}
}

func (ol *OrderListener) HandleExpirationComplete(data []byte) error {
	var event model.ExpirationCompleteEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal expiration complete event")
		return err
	}

	log.Printf("Processing ExpirationCompleteEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateOrderRequest := new(model.UpdateOrderStatusRequest)
	updateOrderRequest.ID = event.OrderID
	updateOrderRequest.Status = domain.Cancelled

	if err := ol.OrderUsecase.UpdateStatus(ctx, updateOrderRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to update order status")
		return err
	}

	return nil
}

func (ol *OrderListener) HandlePaymentCreated(data []byte) error {
	var event model.PaymentCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal payment created event")
		return err
	}

	log.Printf("Processing OrderCancelledEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateOrderRequest := new(model.UpdateOrderStatusRequest)
	updateOrderRequest.ID = event.OrderID
	updateOrderRequest.Status = domain.Complete

	if err := ol.OrderUsecase.UpdateStatus(ctx, updateOrderRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to update order status")
		return err
	}

	return nil
}

func (ol *OrderListener) HandleTicketCreated(data []byte) error {
	var event model.TicketCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal ticket created event")
		return err
	}

	log.Printf("Processing TicketCreatedEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketRequest := new(model.TicketRequest)
	ticketRequest.ID = event.ID
	ticketRequest.Title = event.Title
	ticketRequest.Price = event.Price

	if err := ol.TicketUsecase.Create(ctx, ticketRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to create ticket")
		return err
	}

	return nil
}

func (ol *OrderListener) HandleTicketUpdated(data []byte) error {
	var event model.TicketUpdatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal ticket updated event")
		return err
	}

	log.Printf("Processing TicketUpdatedEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketRequest := new(model.TicketRequest)
	ticketRequest.ID = event.ID
	ticketRequest.Title = event.Title
	ticketRequest.Price = event.Price

	if err := ol.TicketUsecase.Update(ctx, ticketRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to create ticket")
		return err
	}

	return nil
}

func (ol *OrderListener) Listen() {
	expirationCompleteListener := &event.Listener{
		Subject:       domain.ExpirationComplete,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleExpirationComplete,
	}

	paymentCreatedListener := &event.Listener{
		Subject:       domain.PaymentCreated,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandlePaymentCreated,
	}

	ticketCreatedListener := &event.Listener{
		Subject:       domain.TicketCreated,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleTicketCreated,
	}

	ticketUpdatedListener := &event.Listener{
		Subject:       domain.TicketUpdated,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleTicketUpdated,
	}

	go expirationCompleteListener.Listener()
	go paymentCreatedListener.Listener()
	go ticketCreatedListener.Listener()
	go ticketUpdatedListener.Listener()
}
