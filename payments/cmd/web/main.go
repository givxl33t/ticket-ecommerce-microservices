package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"ticketing/payments/config"
	"ticketing/payments/internal/infrastructure"
	"ticketing/payments/internal/infrastructure/middleware"
	"ticketing/payments/internal/interface/http/handler"
	"ticketing/payments/internal/interface/http/route"
	"ticketing/payments/internal/interface/listener"
	"ticketing/payments/internal/publisher"
	"ticketing/payments/internal/repository"
	"ticketing/payments/internal/usecase"
)

// @title Payments Service
// @version 1.0
// @description Payments Service HTTP API Docs
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
	stripeKey := config.GetString("STRIPE_KEY")
	db := infrastructure.NewGorm(config)
	natsConn := infrastructure.NewNATS(config)
	logger := infrastructure.NewLogger(config)
	validate := infrastructure.NewValidator(config)

	// Close NATS singleton connection
	defer infrastructure.CloseNATS()

	// Repository, use case, and handler setup
	orderRepository := repository.NewOrderRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)
	paymentPublisher := publisher.NewPaymentPublisher(natsConn)
	paymentGateway := infrastructure.NewStripe(stripeKey)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository, paymentPublisher, paymentGateway, orderRepository, logger, validate, config)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase, logger)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, logger, validate, config)

	// HTTP routes
	authMiddleware := middleware.NewAuth(logger, config)
	route.RegisterRoute(app, paymentHandler, authMiddleware)

	// // Listeners
	orderListener := listener.NewOrderListener(orderUsecase, natsConn, logger)
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
