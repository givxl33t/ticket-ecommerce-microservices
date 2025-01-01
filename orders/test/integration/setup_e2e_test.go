package integration

import (
	"testing"

	"ticketing/orders/config"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/infrastructure"
	"ticketing/orders/internal/infrastructure/middleware"
	"ticketing/orders/internal/interface/http/handler"
	"ticketing/orders/internal/interface/http/route"
	"ticketing/orders/internal/publisher"
	"ticketing/orders/internal/repository"
	"ticketing/orders/internal/usecase"

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
	Config           *viper.Viper
	App              *fiber.App
	DB               *gorm.DB
	NATS             *nats.Conn
	Log              *logrus.Logger
	Validate         *validator.Validate
	OrderRepository  repository.OrderRepository
	TicketRepository repository.TicketRepository
	OrderPublisher   publisher.OrderPublisher
	OrderUsecase     usecase.OrderUsecase
	TicketUsecase    usecase.TicketUsecase
	OrderHandler     *handler.OrderHandler
	AuthMiddleware   fiber.Handler
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
	s.TicketRepository = repository.NewTicketRepository(s.DB)
	s.OrderPublisher = publisher.NewOrderPublisher(s.NATS)
	s.OrderUsecase = usecase.NewOrderUsecase(s.OrderRepository, s.TicketRepository, s.OrderPublisher, s.Log, s.Validate, s.Config)
	s.TicketUsecase = usecase.NewTicketUsecase(s.TicketRepository, s.Log, s.Validate, s.Config)
	s.OrderHandler = handler.NewOrderHandler(s.OrderUsecase, s.Log)
	s.AuthMiddleware = middleware.NewAuth(s.Log, s.Config)
	route.RegisterRoute(s.App, s.OrderHandler, s.AuthMiddleware)
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(s.DB.Migrator().AutoMigrate(&domain.Ticket{}, &domain.Order{}))
}

func (s *e2eTestSuite) TearDownTest() {
	s.Require().NoError(s.DB.Migrator().DropTable("tickets", "orders"))
}
