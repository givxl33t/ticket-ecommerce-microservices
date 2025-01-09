package integration

import (
	"testing"

	"ticketing/payments/config"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/infrastructure"
	"ticketing/payments/internal/infrastructure/middleware"
	"ticketing/payments/internal/interface/http/handler"
	"ticketing/payments/internal/interface/http/route"
	"ticketing/payments/internal/publisher"
	"ticketing/payments/internal/repository"
	"ticketing/payments/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	Config            *viper.Viper
	App               *fiber.App
	DB                *gorm.DB
	NATS              *nats.Conn
	Log               *logrus.Logger
	Validate          *validator.Validate
	OrderRepository   repository.OrderRepository
	PaymentRepository repository.PaymentRepository
	PaymentPublisher  publisher.PaymentPublisher
	PaymentGateway    infrastructure.PaymentGateway
	PaymentUsecase    usecase.PaymentUsecase
	PaymentHandler    *handler.PaymentHandler
	OrderUsecase      usecase.OrderUsecase
	AuthMiddleware    fiber.Handler
}

func TestE2eSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	s.Config = config.New()
	s.DB = infrastructure.NewGorm(s.Config)
	s.NATS = infrastructure.NewNATS(s.Config)
	s.Log = infrastructure.NewLogger(s.Config)
	s.App = infrastructure.NewFiber(s.Config)
	s.Validate = infrastructure.NewValidator(s.Config)
	s.OrderRepository = repository.NewOrderRepository(s.DB)
	s.PaymentRepository = repository.NewPaymentRepository(s.DB)
	s.PaymentPublisher = publisher.NewPaymentPublisher(s.NATS)
	s.PaymentGateway = infrastructure.NewStripe(s.Config)
	s.OrderUsecase = usecase.NewOrderUsecase(s.OrderRepository, s.Log, s.Validate, s.Config)
	s.PaymentUsecase = usecase.NewPaymentUsecase(s.PaymentRepository, s.PaymentPublisher, s.PaymentGateway, s.OrderRepository, s.Log, s.Validate, s.Config)
	s.PaymentHandler = handler.NewPaymentHandler(s.PaymentUsecase, s.Log)
	s.AuthMiddleware = middleware.NewAuth(s.Log, s.Config)
	route.RegisterRoute(s.App, s.PaymentHandler, s.AuthMiddleware)
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(s.DB.Migrator().AutoMigrate(&domain.Payment{}, &domain.Order{}))
}

func (s *e2eTestSuite) TearDownTest() {
	s.Require().NoError(s.DB.Migrator().DropTable("orders", "payments"))
}
