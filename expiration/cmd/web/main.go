package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"ticketing/expiration/config"
	"ticketing/expiration/internal/domain"
	"ticketing/expiration/internal/infrastructure"
	"ticketing/expiration/internal/interface/listener"
	"ticketing/expiration/internal/interface/worker"
	"ticketing/expiration/internal/publisher"
	"ticketing/expiration/internal/usecase"

	"github.com/hibiken/asynq"
)

func main() {
	config := config.New()
	natsConn := infrastructure.NewNATS(config)
	asynqClient := infrastructure.NewAsynqClient(config)
	asynqServer := infrastructure.NewAsynqServer(config)
	expirationPublisher := publisher.NewExpirationPublisher(natsConn)
	logger := infrastructure.NewLogger(config)

	// Close NATS singleton connection
	defer infrastructure.CloseNATS()

	// use case, and handler setup
	expirationUsecase := usecase.NewExpirationUsecase(asynqClient, logger)
	asynqWorker := worker.NewAsynqWorker(expirationPublisher, logger)

	// Listeners
	expirationListener := listener.NewExpirationListener(expirationUsecase, natsConn, logger)
	expirationListener.Listen()

	// Set up Task Handlers
	mux := asynq.NewServeMux()
	mux.HandleFunc(domain.OrderExpirationTask, asynqWorker.HandleOrderExpirationTask)

	go func() {
		log.Println("Starting Asynq Worker")
		if err := asynqServer.Run(mux); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch // blocks the main thread until an interrupt is received

	// cleanup tasks go here
	asynqServer.Stop()
	_ = asynqClient.Close()

	fmt.Println("App shuts down successfully!")

}
