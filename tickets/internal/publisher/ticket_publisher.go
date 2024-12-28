package publisher

import (
	"encoding/json"
	"log"
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/model"

	"github.com/nats-io/nats.go"
)

type TicketPublisher interface {
	Created(ticket *domain.Ticket) error
	Updated(ticket *domain.Ticket) error
}

type TicketPublisherImpl struct {
	NatsConn *nats.Conn
}

func NewTicketPublisher(natsConn *nats.Conn) TicketPublisher {
	return &TicketPublisherImpl{
		NatsConn: natsConn,
	}
}

func (p *TicketPublisherImpl) Created(ticket *domain.Ticket) error {
	message := model.TicketCreatedEvent{
		ID:      ticket.ID,
		Title:   ticket.Title,
		Price:   ticket.Price,
		UserID:  ticket.UserID,
		OrderID: nil,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.TicketCreated, data)
	if err != nil {
		return err
	}

	log.Printf("Published TicketCreatedEvent: %v\n", message)

	return nil
}

func (p *TicketPublisherImpl) Updated(ticket *domain.Ticket) error {
	message := model.TicketUpdatedEvent{
		ID:      ticket.ID,
		Title:   ticket.Title,
		Price:   ticket.Price,
		UserID:  ticket.UserID,
		OrderID: ticket.OrderID,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.TicketUpdated, data)
	if err != nil {
		return err
	}

	log.Printf("Published TicketUpdatedEvent: %v\n", message)

	return nil
}
