package paymentgo

import (
	canopusgo "github.com/alterraacademy/canopus-gopkg"
)

type Payment struct {
	Canopus *canopusgo.Canopus
}

func CreateService() *Payment {

	return &Payment{}
}

func (pay *Payment) InitCanopus(init CanopusService) error {
	cano, err := canopusgo.CreateService(canopusgo.InitService{
		Type:        init.CanopusType,
		MerchantKey: init.MerchantKey,
		MerchantPem: init.MerchantPem,
		TimeOut:     init.TimeOut,
		MerchantID:  init.MerchantID,
		Secret:      init.Secret,
	})

	if err != nil {
		return err
	}

	pay.Canopus = cano

	return nil
}

func (pay *Payment) GetAvailableMethod(amount float64) ([]PaymentMethod, error) {
	var result []PaymentMethod

	// Get from Canopus
	if pay.Canopus != nil {
		canoPaymentMethods, err := pay.Canopus.GetAvailableMethod(amount)
		if err != nil {
			return []PaymentMethod{}, err
		}
		for _, paymentMethod := range canoPaymentMethods {
			result = append(result, PaymentMethodFromCanopus(paymentMethod))
		}
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
