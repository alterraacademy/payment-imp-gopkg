package paymentgo

import (
	"errors"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CartPayloadToMidtrans(cart CartPayload) *coreapi.ChargeReq {

	chargeReq := coreapi.ChargeReq{
		PaymentType: coreapi.CoreapiPaymentType(cart.CartDetail.Payment.Type),
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  cart.CartDetail.ID,
			GrossAmt: int64(cart.CartDetail.Amount),
		},
		Items: MidtransItemDetails(cart.ItemDetails),
		CustomerDetails: &midtrans.CustomerDetails{
			FName: cart.CustomerDetail.FirstName,
			LName: cart.CustomerDetail.LastName,
			Email: cart.CustomerDetail.Email,
			Phone: cart.CustomerDetail.Phone,
			BillAddr: &midtrans.CustomerAddress{
				FName:    cart.CustomerDetail.BillingAddress.FirstName,
				LName:    cart.CustomerDetail.BillingAddress.LastName,
				Address:  cart.CustomerDetail.BillingAddress.Address,
				City:     cart.CustomerDetail.BillingAddress.City,
				Postcode: cart.CustomerDetail.BillingAddress.PostalCode,
				Phone:    cart.CustomerDetail.BillingAddress.Phone,
			},
			ShipAddr: &midtrans.CustomerAddress{
				FName:    cart.CustomerDetail.ShippingAddress.FirstName,
				LName:    cart.CustomerDetail.ShippingAddress.LastName,
				Address:  cart.CustomerDetail.ShippingAddress.Address,
				City:     cart.CustomerDetail.ShippingAddress.City,
				Postcode: cart.CustomerDetail.ShippingAddress.PostalCode,
				Phone:    cart.CustomerDetail.ShippingAddress.Phone,
			},
		},
	}

	switch chargeReq.PaymentType {
	case coreapi.PaymentTypeGopay:
		chargeReq.Gopay = &coreapi.GopayDetails{
			EnableCallback: true,
			CallbackUrl:    cart.Callback.Return,
		}
	case coreapi.PaymentTypeShopeepay:
		chargeReq.ShopeePay = &coreapi.ShopeePayDetails{
			CallbackUrl: cart.Callback.Return,
		}
	case coreapi.PaymentTypeBankTransfer:
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(cart.CartDetail.Payment.Key),
		}
	default:
		return &coreapi.ChargeReq{}
	}

	return &chargeReq
}

func MidtransItemDetails(items []CartPayloadItemDetail) *[]midtrans.ItemDetails {
	var mItems []midtrans.ItemDetails

	for _, val := range items {
		tmp := midtrans.ItemDetails{
			ID:    val.SKU,
			Price: int64(val.Price),
			Qty:   int32(val.Quantity),
			Name:  val.Name,
		}
		mItems = append(mItems, tmp)

	}

	return &mItems
}

func CartResponseFromMidtrans(resp coreapi.ChargeResponse) (CartResponse, error) {

	if !strings.Contains(resp.StatusCode, "20") {
		return CartResponse{}, errors.New(resp.StatusMessage)
	}

	var link string = ""
	if len(resp.Actions) > 0 {
		for _, val := range resp.Actions {
			if val.Name == "deeplink-redirect" {
				link = val.URL
			}
		}
	}

	var bank string = ""
	var number string = ""
	if len(resp.VaNumbers) > 0 {
		bank = resp.VaNumbers[0].Bank
		number = resp.VaNumbers[0].VANumber
	}

	if strings.Contains(strings.ToLower(resp.StatusMessage), "permata") {
		bank = "permata"
		number = resp.PermataVaNumber
	}

	return CartResponse{
		CartID: resp.TransactionID,
		Link:   link,
		Amount: resp.GrossAmount,
		Bank:   bank,
		Number: number,
	}, nil
}
