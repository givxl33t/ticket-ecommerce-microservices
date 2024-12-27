package domain

const (
	// Ticket-related events
	TicketCreated = "ticket:created"
	TicketUpdated = "ticket:updated"

	// Order-related events
	OrderCreated   = "order:created"
	OrderCancelled = "order:cancelled"

	// System-related events
	ExpirationComplete = "expiration:complete"
	PaymentCreated     = "payment:created"
)
