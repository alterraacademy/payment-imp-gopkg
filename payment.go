package paymentgo

import (
	canopusgo "github.com/alterraacademy/canopus-gopkg"
	canopusConfig "github.com/alterraacademy/canopus-gopkg/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentContract interface {
	GetAvailableMethod(amount float64) ([]PaymentMethod, error)
	CreateCart(cart CartPayload, method PaymentMethod) (CartResponse, error)
}

func InitPayment(paymentType PaymentType, init CanopusService, mTransServ MidtransService) PaymentContract {
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

	if paymentType == Midtrans {
		mTransClient := coreapi.Client{}
		environment := midtrans.Production
		if mTransServ.Sandbox {
			environment = midtrans.Sandbox
		}
		mTransClient.New(mTransServ.ServerKey, environment)

		return &MidtransPayment{
			Client: mTransClient,
		}
	}

	return nil
}
