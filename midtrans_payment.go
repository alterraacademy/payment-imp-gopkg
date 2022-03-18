package paymentgo

import (
	"errors"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransPayment struct {
	Client coreapi.Client
}

func (mtrans *MidtransPayment) GetAvailableMethod(amount float64) ([]PaymentMethod, error) {

	return []PaymentMethod{
		{
			Type: string(coreapi.PaymentTypeGopay),
			Name: "Gopay",
			Key:  string(coreapi.PaymentTypeGopay),
		},
		{
			Type: string(coreapi.PaymentTypeShopeepay),
			Name: "Shopeepay",
			Key:  string(coreapi.PaymentTypeShopeepay),
		},
		{
			Type: string(coreapi.PaymentTypeBankTransfer),
			Name: "VA Permata",
			Key:  string(midtrans.BankPermata),
		},
		{
			Type: string(coreapi.PaymentTypeBankTransfer),
			Name: "VA BCA",
			Key:  string(midtrans.BankBca),
		},
		{
			Type: string(coreapi.PaymentTypeBankTransfer),
			Name: "VA BNI",
			Key:  string(midtrans.BankBni),
		},
		{
			Type: string(coreapi.PaymentTypeBankTransfer),
			Name: "VA BRI",
			Key:  string(midtrans.BankBri),
		},
		{
			Type: string(coreapi.PaymentTypeCreditCard),
			Name: "Credit Card",
			Key:  string(coreapi.PaymentTypeCreditCard),
		},
	}, nil
}

func (mtrans *MidtransPayment) CreateCart(cart CartPayload, method PaymentMethod) (CartResponse, error) {
	if mtrans.Client == (coreapi.Client{}) {
		return CartResponse{}, ErrNoMidtransService
	}
	mtrans.Client.Options.SetPaymentOverrideNotification(cart.Callback.Hook)
	result, err := mtrans.Client.ChargeTransaction(CartPayloadToMidtrans(cart))
	if err != nil {
		return CartResponse{}, err.RawError
	}

	return CartResponseFromMidtrans(*result)
}

func (mtrans *MidtransPayment) CardRegister(cc CreditCard) (string, error) {
	if mtrans.Client == (coreapi.Client{}) {
		return "", ErrNoMidtransService
	}

	cvv := fmt.Sprintf("%d", cc.CVV)

	resp, err := mtrans.Client.CardToken(cc.CardNumber, cc.ExpMonth, cc.ExpYear, cvv, midtrans.ClientKey)
	if err != nil {
		return "", errors.New("error get card token : " + err.GetMessage())
	}
	return resp.TokenID, nil
}
