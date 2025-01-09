package integration

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"ticketing/payments/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

func (s *e2eTestSuite) TestPaymentCreatedSuccess() {
	s.CreateDummyOrder(
		1,
		"created",
		"user-1",
		1000,
	)

	requestBody := &model.PaymentRequest{
		UserID:  "user-1",
		OrderID: 1,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/payments", strings.NewReader(string(bodyJSON)))
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

	responseBody := new(model.PaymentResponse)
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)
}

func (s *e2eTestSuite) TestOrderCreatedFailedValidation() {
	requestBody := &model.CreateOrderRequest{
		UserID: "",
		Price:  0,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/payments", strings.NewReader(string(bodyJSON)))
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
		"id":    "user-1",
		"email": "JlKk2@example.com",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.Config.GetString("JWT_SECRET_KEY")))
	s.Assert().NoError(err)

	return token
}

// create a dummy order data
func (s *e2eTestSuite) CreateDummyOrder(ID int, Status string, UserID string, Price int) {
	order := &model.CreateOrderRequest{
		ID:     int32(ID),
		Status: Status,
		UserID: UserID,
		Price:  int64(Price),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.OrderUsecase.Create(ctx, order); err != nil {
		s.Log.WithError(err).Error("failed to create order")
	}
}
