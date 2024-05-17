package services

import (
	"log"
	"restaurant-payment/config"
	"restaurant-payment/models"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

type StripeService struct {
	config config.Config
}

func NewStripeService(config config.Config) *StripeService {
	stripe.Key = config.StripeSecretKey
	return &StripeService{config: config}
}

func (s *StripeService) Charge(order models.Order, token string) (*stripe.Charge, error) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(order.Amount),
		Currency:    stripe.String(order.Currency),
		Description: stripe.String("Restaurant Payment"),
	}
	params.SetSource(token)

	charge, err := charge.New(params)
	if err != nil {
		log.Printf("Stripe charge failed: %v", err)
		return nil, err
	}
	return charge, nil
}
