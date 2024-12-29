package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"ticketing/tickets/config"
	"ticketing/tickets/internal/infrastructure"
	"ticketing/tickets/internal/infrastructure/middleware"
	"ticketing/tickets/internal/interface/http/handler"
	"ticketing/tickets/internal/interface/http/route"
	"ticketing/tickets/internal/interface/listener"
	"ticketing/tickets/internal/publisher"
	"ticketing/tickets/internal/repository"
	"ticketing/tickets/internal/usecase"
)

// @title Tickets Service
// @version 1.0
// @description Tickets Service HTTP API Docs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apikey Session
// @in header
// @name Cookie
func main() {
	config := config.New()

	// Infrastructure setup
	app := infrastructure.NewFiber(config)
	port := config.Get("APP_PORT")
	db := infrastructure.NewGorm(config)
	natsConn := infrastructure.NewNATS(config)
	logger := infrastructure.NewLogger(config)
	validate := infrastructure.NewValidator(config)

	// Close NATS singleton connection
	defer infrastructure.CloseNATS()

	// Repository, use case, and handler setup
	ticketRepository := repository.NewTicketRepository(db)
	ticketPublisher := publisher.NewTicketPublisher(natsConn)
	TicketUsecase := usecase.NewTicketUsecase(ticketRepository, ticketPublisher, logger, validate, config)
	ticketHandler := handler.NewTicketHandler(TicketUsecase, logger)

	// HTTP routes
	authMiddleware := middleware.NewAuth(logger, config)
	route.RegisterRoute(app, ticketHandler, authMiddleware)

	// Listeners
	orderListener := listener.NewOrderListener(TicketUsecase, natsConn, logger)
	orderListener.Listen()

	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
			panic(fmt.Errorf("error running app : %+v", err.Error()))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch // blocks the main thread until an interrupt is received

	// cleanup tasks go here
	_ = app.Shutdown()

	fmt.Println("App shuts down successfully!")
}
