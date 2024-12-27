package publisher

import (
	"ticketing/tickets/internal/common/event"

	"github.com/nats-io/nats.go"
)

type TicketUpdatedPublisher struct {
	*event.Publisher
}

func NewTicketUpdatedPublisher(conn *nats.Conn) *TicketUpdatedPublisher {
	return &TicketUpdatedPublisher{
		Publisher: &event.Publisher{
			Subject:  event.TicketUpdated,
			NatsConn: conn,
		},
	}
}
