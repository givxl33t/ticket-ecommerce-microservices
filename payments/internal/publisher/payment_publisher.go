package publisher

import (
	"encoding/json"
	"log"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/model"

	"github.com/nats-io/nats.go"
)

type PaymentPublisher interface {
	Created(order *domain.Payment) error
}

type PaymentPublisherImpl struct {
	NatsConn *nats.Conn
}

func NewPaymentPublisher(natsConn *nats.Conn) PaymentPublisher {
	return &PaymentPublisherImpl{
		NatsConn: natsConn,
	}
}

func (p *PaymentPublisherImpl) Created(order *domain.Payment) error {
	message := model.PaymentCreatedEvent{
		ID:       order.ID,
		OrderID:  order.OrderID,
		StripeID: order.StripeID,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.PaymentCreated, data)
	if err != nil {
		return err
	}

	log.Printf("Published event on subject: %s", domain.PaymentCreated)

	return nil
}
