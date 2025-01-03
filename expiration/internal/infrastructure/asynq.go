package infrastructure

import (
	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
)

func NewAsynqClient(config *viper.Viper) *asynq.Client {
	host := config.GetString("REDIS_HOST")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: host})
	return client
}

func NewAsynqServer(config *viper.Viper) *asynq.Server {
	host := config.GetString("REDIS_HOST")
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: host},
		asynq.Config{
			Concurrency: 10,
		})

	return server
}
