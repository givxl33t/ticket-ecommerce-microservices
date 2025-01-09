package event

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Listener struct {
	Subject       string
	QueueGroup    string
	NatsConn      *nats.Conn
	AckWait       time.Duration
	OnMessageFunc func(data []byte) error
}

func (l *Listener) Listener() {
	_, err := l.NatsConn.QueueSubscribe(l.Subject, l.QueueGroup, func(msg *nats.Msg) {
		log.Printf("Message received: Subject: %s, Queue: %s\n", l.Subject, l.QueueGroup)

		if err := l.OnMessageFunc(msg.Data); err != nil {
			log.Printf("Error processing message: %v\n", err)
			return
		}

		// Acknowledge the message if processing succeds (optional, depending on implementation)
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Printf("Listening on subject: %s, queue: %s\n", l.Subject, l.QueueGroup)

	// Do not defer unsubscribe or else after a successfull listen it would unsubscribe immediately
}
