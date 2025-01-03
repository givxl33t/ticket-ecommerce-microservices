package publisher

import (
	"encoding/json"
	"log"
	"ticketing/expiration/internal/domain"
	"ticketing/expiration/internal/model"

	"github.com/nats-io/nats.go"
)

type ExpirationPublisher interface {
	Expired(orderId int32) error
}

type ExpirationPublisherImpl struct {
	NatsConn *nats.Conn
}

func NewExpirationPublisher(natsConn *nats.Conn) ExpirationPublisher {
	return &ExpirationPublisherImpl{
		NatsConn: natsConn,
	}
}

func (p *ExpirationPublisherImpl) Expired(orderId int32) error {
	message := model.OrderExpirationPayload{
		OrderID: orderId,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.ExpirationComplete, data)
	if err != nil {
		return err
	}

	log.Printf("Published event on subject: %s", domain.OrderCreated)

	return nil
}
