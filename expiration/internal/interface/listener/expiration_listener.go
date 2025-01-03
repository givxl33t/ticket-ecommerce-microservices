package listener

import (
	"context"
	"encoding/json"
	"log"
	"ticketing/expiration/internal/common/event"
	"ticketing/expiration/internal/domain"
	"ticketing/expiration/internal/model"
	"ticketing/expiration/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	QueueGroupName = "expiration-service"
)

type ExpirationListener struct {
	ExpirationUsecase usecase.ExpirationUsecase
	NatsConn          *nats.Conn
	Logger            *logrus.Logger
}

func NewExpirationListener(expirationUsecase usecase.ExpirationUsecase, conn *nats.Conn, log *logrus.Logger) *ExpirationListener {
	return &ExpirationListener{
		ExpirationUsecase: expirationUsecase,
		NatsConn:          conn,
		Logger:            log,
	}
}

func (ol *ExpirationListener) HandleOrderCreated(data []byte) error {
	var event model.OrderCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		ol.Logger.WithError(err).Error("failed unmarshal ticket updated event")
		return err
	}

	log.Printf("Processing OrderCreatedEvent: %v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	orderPayload := new(model.OrderExpirationPayload)
	orderPayload.OrderID = event.ID
	orderPayload.Delay = event.ExpiresAt - time.Now().Unix()

	if err := ol.ExpirationUsecase.ScheduleExpiration(ctx, orderPayload); err != nil {
		ol.Logger.WithError(err).Error("failed to schedule expiration")
		return err
	}

	return nil
}

func (ol *ExpirationListener) Listen() {
	ticketUpdatedListener := &event.Listener{
		Subject:       domain.TicketUpdated,
		QueueGroup:    QueueGroupName,
		NatsConn:      ol.NatsConn,
		OnMessageFunc: ol.HandleOrderCreated,
	}

	go ticketUpdatedListener.Listener()
}
