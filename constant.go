package paymentgo

import "errors"

type PaymentType string

const (
	Canopus  PaymentType = "canopus"
	Midtrans PaymentType = "midtrans"

	CanopusSNAP = "snap"
	CanopusAPI  = "api"
)

var (
	ErrNoCanopusService  = errors.New("no canopus service yet, please init first")
	ErrNoMidtransService = errors.New("no midtrans service yet, please init first")
)
