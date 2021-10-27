package paymentgo

import "errors"

type PaymentType uint8

const (
	_ PaymentType = iota
	Canopus

	CanopusSNAP = "snap"
	CanopusAPI  = "api"
)

var (
	ErrNoCanopusService = errors.New("no canopus service yet, please init first")
)
