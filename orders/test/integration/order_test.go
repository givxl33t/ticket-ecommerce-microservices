package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"ticketing/tickets/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

func (s *e2eTestSuite) TestTicketCreatedSuccess() {
	requestBody := &model.CreateTicketRequest{
		Title: "concert",
		Price: 39,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/tickets", strings.NewReader(string(bodyJSON)))
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

	responseBody := new(model.TicketResponse)
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().Equal(requestBody.Title, responseBody.Title)
}

func (s *e2eTestSuite) TestTicketCreatedFailedValidation() {
	requestBody := &model.CreateTicketRequest{
		Title: "",
		Price: 0,
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/tickets", strings.NewReader(string(bodyJSON)))
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
