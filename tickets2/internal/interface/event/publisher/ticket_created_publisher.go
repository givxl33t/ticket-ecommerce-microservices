package publisher

import (
	"ticketing/tickets/internal/common/event"

	"github.com/nats-io/nats.go"
)

type TicketCreatedPublisher struct {
	*event.Publisher
}

func NewTicketCreatedPublisher(conn *nats.Conn) *TicketCreatedPublisher {
	return &TicketCreatedPublisher{
		Publisher: &event.Publisher{
			Subject:  event.TicketCreated,
			NatsConn: conn,
		},
	}
}
