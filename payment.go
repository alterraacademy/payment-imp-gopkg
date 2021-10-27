package paymentgo

import (
	canopusgo "github.com/alterraacademy/canopus-gopkg"
	canopusConfig "github.com/alterraacademy/canopus-gopkg/config"
)

type PaymentContract interface {
	GetAvailableMethod(amount float64) ([]PaymentMethod, error)
	CreateCart(cart CartPayload, method PaymentMethod) (CartResponse, error)
}

func InitPayment(paymentType PaymentType, init CanopusService) PaymentContract {
	if paymentType == Canopus {
		canopusConfig.DefaultAPIType = canopusConfig.API
		if init.CanopusType == CanopusSNAP {
			canopusConfig.DefaultAPIType = canopusConfig.SNAP
		}

		paymentClient := canopusgo.NewAPICLient(&canopusgo.ConfigOptions{
			MerchantKey: init.MerchantKey,
			MerchantPem: init.MerchantPem,
			MerchantID:  init.MerchantID,
			Secret:      init.Secret,
			Timeout:     int(init.TimeOut),
		})

		return &Payment{
			Canopus: paymentClient,
		}

	}

	return nil
}
