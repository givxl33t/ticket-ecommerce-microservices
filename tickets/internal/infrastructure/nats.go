package infrastructure

import (
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

var (
	natsInstance *nats.Conn
	once         sync.Once
)

func NewNATS(config *viper.Viper) (*nats.Conn, error) {
	var err error

	once.Do(func() {
		natsURL := config.GetString("NATS_URL")

		natsInstance, err = nats.Connect(natsURL)
	})

	return natsInstance, err
}

func CloseNATS() {
	if natsInstance != nil {
		natsInstance.Close()
	}
}
