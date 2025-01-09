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

	// singleton connection
	once.Do(func() {
		natsURL := config.GetString("NATS_URL")
		maxRetries := 10

		// test retry logic due to its heavy reliance
		natsInstance, err = nats.Connect(natsURL, nats.RetryOnFailedConnect(true), nats.MaxReconnects(maxRetries))

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
