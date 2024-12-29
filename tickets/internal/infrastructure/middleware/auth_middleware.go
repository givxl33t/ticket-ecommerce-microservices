package middleware

import (
	"ticketing/tickets/internal/common/exception"
	"ticketing/tickets/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewAuth(logger *logrus.Logger, config *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("session")

		if accessToken == "" {
			return exception.ErrUserUnauthorized
		}

		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			logger.WithError(err).Error("user unauthorized")
			return exception.ErrUserUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return exception.ErrUserUnauthorized
		}

		auth := &model.AccessTokenPayload{
			Email: claims["email"].(string),
			ID:    claims["id"].(string),
		}

		c.Locals("auth", auth)

		return c.Next()
	}
}
