package infrastructure

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

var (
	natsInstance *nats.Conn
	once         sync.Once
)

func NewNATS(config *viper.Viper) *nats.Conn {
	var err error

	once.Do(func() {
		natsURL := config.GetString("NATS_URL")

		natsInstance, err = nats.Connect(natsURL)

		if err != nil {
			panic(fmt.Errorf("failed to connect to NATS: %v", err))
		}
	})

	return natsInstance
}

func CloseNATS() {
	if natsInstance != nil {
		natsInstance.Close()
	}
}
