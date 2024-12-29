package integration

import (
	"testing"

	"ticketing/tickets/config"
	"ticketing/tickets/internal/domain"
	"ticketing/tickets/internal/infrastructure"
	"ticketing/tickets/internal/infrastructure/middleware"
	"ticketing/tickets/internal/interface/http/handler"
	"ticketing/tickets/internal/interface/http/route"
	"ticketing/tickets/internal/publisher"
	"ticketing/tickets/internal/repository"
	"ticketing/tickets/internal/usecase"

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
	TicketRepository repository.TicketRepository
	TicketPublisher  publisher.TicketPublisher
	TicketUsecase    usecase.TicketUsecase
	TicketHandler    *handler.TicketHandler
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
	s.TicketRepository = repository.NewTicketRepository(s.DB)
	s.TicketPublisher = publisher.NewTicketPublisher(s.NATS)
	s.TicketUsecase = usecase.NewTicketUsecase(s.TicketRepository, s.TicketPublisher, s.Log, s.Validate, s.Config)
	s.TicketHandler = handler.NewTicketHandler(s.TicketUsecase, s.Log)
	s.AuthMiddleware = middleware.NewAuth(s.Log, s.Config)
	route.RegisterRoute(s.App, s.TicketHandler, s.AuthMiddleware)
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(s.DB.Migrator().AutoMigrate(&domain.Ticket{}))
}

func (s *e2eTestSuite) TearDownTest() {
	s.Require().NoError(s.DB.Migrator().DropTable("tickets"))
}
