package infrastructure

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

// make it a singleton instance okok
type NatsWrapper struct {
	conn  *nats.Conn
	mu    sync.Mutex
	ready bool
}

var instance *NatsWrapper

func GetNatsWrapper() *NatsWrapper {
	if instance == nil {
		instance = &NatsWrapper{}
	}
	return instance
}

func (nw *NatsWrapper) Connect(config *viper.Viper) error {
	url := config.GetString("NATS_URL")

	nw.mu.Lock()
	defer nw.mu.Unlock()
	if nw.ready {
		return nil
	}
	conn, err := nats.Connect(url)
	if err != nil {
		return err
	}
	nw.conn = conn
	nw.ready = true

	log.Println("Connected to NATS")
	return nil
}

func (nw *NatsWrapper) Conn() *nats.Conn {
	if !nw.ready {
		panic("Cannot access NATS connection before connecting")
	}
	return nw.conn
}
