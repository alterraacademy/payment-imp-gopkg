package paymentgo

import (
	canopusgo "github.com/alterraacademy/canopus-gopkg"
)

type Payment struct {
	Canopus canopusgo.ClientMethod
}

func (pay *Payment) GetAvailableMethod(amount float64) ([]PaymentMethod, error) {
	var result []PaymentMethod
	if pay.Canopus == nil {
		return result, ErrNoCanopusService
	}
	// Get from Canopus
	canoPaymentMethods, err := pay.Canopus.GetAvailableMethod(amount)
	if err != nil {
		return []PaymentMethod{}, err
	}
	for _, paymentMethod := range canoPaymentMethods {
		result = append(result, PaymentMethodFromCanopus(paymentMethod))
	}

	return result, nil
}

func (pay *Payment) CreateCart(cart CartPayload, method PaymentMethod) (CartResponse, error) {
	if pay.Canopus == nil {
		return CartResponse{}, ErrNoCanopusService
	}

	cano, err := pay.Canopus.GenerateCart(CartPayloadToCanopus(cart), PaymentMethodToCanopus(method))

	if err != nil {
		return CartResponse{}, err
	}

	return CartResponseFromCanopus(cano), nil
}
