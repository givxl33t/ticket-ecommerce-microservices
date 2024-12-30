package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"ticketing/orders/config"
	"ticketing/orders/internal/infrastructure"
	"ticketing/orders/internal/infrastructure/middleware"
	"ticketing/orders/internal/interface/http/handler"
	"ticketing/orders/internal/interface/http/route"
	"ticketing/orders/internal/interface/listener"
	"ticketing/orders/internal/publisher"
	"ticketing/orders/internal/repository"
	"ticketing/orders/internal/usecase"
)

// @title Orders Service
// @version 1.0
// @description Tickets Service HTTP API Docs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
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
	orderRepository := repository.NewOrderRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	orderPublisher := publisher.NewOrderPublisher(natsConn)
	ticketUsecase := usecase.NewTicketUsecase(ticketRepository, logger, validate, config)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, ticketRepository, orderPublisher, logger, validate, config)
	orderHandler := handler.NewOrderHandler(orderUsecase, logger)

	// HTTP routes
	authMiddleware := middleware.NewAuth(logger, config)
	route.RegisterRoute(app, orderHandler, authMiddleware)

	// // Listeners
	orderListener := listener.NewOrderListener(orderUsecase, ticketUsecase, natsConn, logger)
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
