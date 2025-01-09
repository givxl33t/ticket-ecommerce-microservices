package infrastructure

import (
	"fmt"
	"ticketing/payments/internal/domain"

	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type PaymentGateway interface {
	CreatePayment(order *domain.Order) (*stripe.PaymentIntent, error)
}

type PaymentGatewayImpl struct {
	Config *viper.Viper
}

func NewStripe(config *viper.Viper) PaymentGateway {
	return &PaymentGatewayImpl{
		Config: config,
	}
}

func (pi *PaymentGatewayImpl) CreatePayment(order *domain.Order) (*stripe.PaymentIntent, error) {
	stripe.Key = pi.Config.GetString("STRIPE_KEY")
	intentParams := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(order.Price),
		Currency:           stripe.String("usd"),
		Description:        stripe.String(fmt.Sprintf("Ticket for %d", order.ID)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Params: stripe.Params{
			Metadata: map[string]string{
				"order_id": fmt.Sprintf("%d", order.ID),
			},
		},
	}
	paymentIntent, err := paymentintent.New(intentParams)
	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}
