package event

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	Subject  string
	NatsConn *nats.Conn
}

func (p *Publisher) Publish(subject string, data interface{}) error {
	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := p.NatsConn.Publish(subject, encodedData); err != nil {
		return err
	}

	log.Printf("Event published to subject %v", p.Subject)
	return nil
}
