package listener

import (
	"context"
	"encoding/json"
	"log"
	"ticketing/payments/internal/common/event"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/model"
	"ticketing/payments/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	QueueGroupName = "payments-service"
)

type OrderListener struct {
	OrderUsecase usecase.OrderUsecase
	NatsConn     *nats.Conn
	Logger       *logrus.Logger
}

func NewOrderListener(orderUseCase usecase.OrderUsecase, conn *nats.Conn, log *logrus.Logger) *OrderListener {
	return &OrderListener{
		OrderUsecase: orderUseCase,
		NatsConn:     conn,
		Logger:       log,
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

	createOrderRequest := new(model.CreateOrderRequest)
	createOrderRequest.Status = event.Status
	createOrderRequest.Price = event.Ticket.Price
	createOrderRequest.UserID = event.UserID
	createOrderRequest.ID = event.ID

	if err := ol.OrderUsecase.Create(ctx, createOrderRequest); err != nil {
		ol.Logger.WithError(err).Error("failed to create order")
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

	updateOrderStatusRequest := new(model.UpdateOrderStatusRequest)
	updateOrderStatusRequest.Status = domain.Cancelled
	updateOrderStatusRequest.ID = event.ID

	if err := ol.OrderUsecase.UpdateStatus(ctx, updateOrderStatusRequest); err != nil {
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
