package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ticketing/tickets/internal/common/event"
	"ticketing/tickets/internal/interface/event/publisher"
	"ticketing/tickets/internal/model"
	"ticketing/tickets/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
)

type OrderCancelledEvent struct {
	ID     *int32 `json:"id"`
	Ticket struct {
		ID int32 `json:"id"`
	} `json:"ticket"`
}

type OrderCancelledListener struct {
	*event.Listener
}

func NewOrderCancelledListener(ticketUseCase usecase.TicketUsecase, conn *nats.Conn) *OrderCancelledListener {
	return &OrderCancelledListener{
		Listener: &event.Listener{
			Subject:    event.OrderCancelled,
			QueueGroup: QueueGroupName,
			NatsConn:   conn,
			OnMessageFunc: func(data []byte) error {
				var event OrderCancelledEvent
				publisher := publisher.NewTicketUpdatedPublisher(conn)
				if err := json.Unmarshal(data, &event); err != nil {
					return err
				}

				log.Printf("Processing OrderCancelledEvent: %v\n", event)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				updateTicketOrderRequest := new(model.UpdateTicketOrderRequest)
				updateTicketOrderRequest.ID = event.Ticket.ID

				// Save process to database
				response, err := ticketUseCase.UpdateOrder(ctx, updateTicketOrderRequest)
				if err != nil {
					return fmt.Errorf("failed to update ticket order: %w", err)
				}

				ticketUpdatedEvent := TicketUpdatedEvent{
					ID:      response.ID,
					Title:   response.Title,
					Price:   response.Price,
					UserID:  response.UserID,
					OrderID: response.OrderID,
				}

				if err := publisher.Publish(ticketUpdatedEvent); err != nil {
					return fmt.Errorf("failed to publish TicketUpdatedEvent: %w", err)
				}

				return nil
			},
		},
	}
}
