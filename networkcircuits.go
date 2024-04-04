package glesys

import (
	"context"
)

// NetworkCircuitService provides functions to interact with NetworkCircuits
type NetworkCircuitService struct {
	client clientInterface
}

// NetworkCircuit represents a networkcircuit
type NetworkCircuit struct {
	ID      string                `json:"id"`
	Type    string                `json:"type"`
	Billing NetworkCircuitBilling `json:"billing"`
}

type NetworkCircuitBilling struct {
	Currency string                         `json:"currency"`
	Price    float64                        `json:"price"`
	Discount float64                        `json:"discount"`
	Total    float64                        `json:"total"`
	Details  []NetworkCircuitBillingDetails `json:"details,omitempty"`
}

type NetworkCircuitBillingDetails struct {
	Text                string  `json:"text"`
	PriceBeforeDiscount float64 `json:"subtotalBeforeDiscount"`
	DiscountAmount      float64 `json:"discountAmount"`
	TotalBeforeTax      float64 `json:"totalBeforeTax"`
}

// Details returns detailed information about one NetworkCircuit
func (s *NetworkCircuitService) Details(context context.Context, circuitID string) (*NetworkCircuit, error) {
	data := struct {
		Response struct {
			NetworkCircuit NetworkCircuit
		}
	}{}
	err := s.client.post(context, "networkcircuit/details", &data, circuitID)
	return &data.Response.NetworkCircuit, err
}

// List returns a list of NetworkCircuits available under your account
func (s *NetworkCircuitService) List(context context.Context) (*[]NetworkCircuit, error) {
	data := struct {
		Response struct {
			NetworkCircuits []NetworkCircuit
		}
	}{}

	err := s.client.get(context, "networkcircuit/list", &data)
	return &data.Response.NetworkCircuits, err
}
