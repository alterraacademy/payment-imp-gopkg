package paymentgo

import "errors"

var (
	ErrNoCanopusService = errors.New("no canopus service yet, please init first")
)
