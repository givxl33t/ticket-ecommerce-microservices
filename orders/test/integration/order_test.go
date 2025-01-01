package integration

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"ticketing/orders/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

func (s *e2eTestSuite) TestOrderCreatedSuccess() {
	s.CreateDummyTicket(
		1,
		"Ticket 1",
		1000,
	)

	requestBody := &model.CreateOrderRequest{
		TicketID: 1,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")
	// fake user
	cookie := &http.Cookie{
		Name:  "session",
		Value: s.GenerateUserToken(),
	}
	request.AddCookie(cookie)

	response, err := s.App.Test(request)

	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusCreated, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := new(model.OrderResponse)
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().Equal(requestBody.TicketID, responseBody.Ticket.ID)
}

func (s *e2eTestSuite) TestOrderCreatedFailedValidation() {
	requestBody := &model.CreateOrderRequest{
		TicketID: 0,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")
	// fake user
	cookie := &http.Cookie{
		Name:  "session",
		Value: s.GenerateUserToken(),
	}
	request.AddCookie(cookie)

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

// generate jwt with predefined claims
func (s *e2eTestSuite) GenerateUserToken() string {
	claims := jwt.MapClaims{
		"id":    "asdasddsaads",
		"email": "JlKk2@example.com",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.Config.GetString("JWT_SECRET_KEY")))
	s.Assert().NoError(err)

	return token
}

// create a dummy ticket data
func (s *e2eTestSuite) CreateDummyTicket(ID int, Title string, Price int) {
	ticket := &model.TicketRequest{
		ID:    int32(ID),
		Title: Title,
		Price: int64(Price),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.TicketUsecase.Create(ctx, ticket); err != nil {
		s.Log.WithError(err).Error("failed to create ticket")
	}
}
