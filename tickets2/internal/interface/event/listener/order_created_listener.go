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

type OrderCreatedEvent struct {
	ID        *int32 `json:"id"`
	Status    string `json:"status"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
	Ticket    struct {
		ID    int32 `json:"id"`
		Price int64 `json:"price"`
	} `json:"ticket"`
}

type TicketUpdatedEvent struct {
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Price  int64  `json:"price"`
	UserID string `json:"user_id"`
	// returning the pointer? IDK
	OrderID *int32 `json:"order_id"`
}

type OrderCreatedListener struct {
	*event.Listener
}

func NewOrderCreatedListener(ticketUseCase usecase.TicketUsecase, conn *nats.Conn) *OrderCreatedListener {
	return &OrderCreatedListener{
		Listener: &event.Listener{
			Subject:    event.OrderCreated,
			QueueGroup: QueueGroupName,
			NatsConn:   conn,
			OnMessageFunc: func(data []byte) error {
				var event OrderCreatedEvent
				publisher := publisher.NewTicketUpdatedPublisher(conn)
				if err := json.Unmarshal(data, &event); err != nil {
					return err
				}

				log.Printf("Processing OrderCreatedEvent: %v\n", event)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				updateTicketOrderRequest := new(model.UpdateTicketOrderRequest)
				updateTicketOrderRequest.OrderID = event.ID
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
