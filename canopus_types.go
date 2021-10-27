package paymentgo

import (
	canopusgo "github.com/alterraacademy/canopus-gopkg"
)

func PaymentMethodFromCanopus(cano canopusgo.PaymentMethod) PaymentMethod {
	var result PaymentMethod

	result.From = "canopus"
	result.Name = cano.Name
	result.Type = cano.Type
	result.Key = cano.Key
	result.Logo = cano.Logo
	result.Instruction = cano.Instruction

	return result
}

func PaymentMethodToCanopus(method PaymentMethod) canopusgo.PaymentMethod {
	var result canopusgo.PaymentMethod

	result.Name = method.Name
	result.Type = method.Type
	result.Key = method.Key
	result.Logo = method.Logo
	result.Instruction = method.Instruction

	return result
}

func CartPayloadToCanopus(cart CartPayload) canopusgo.CartPayload {
	var result canopusgo.CartPayload

	result.ItemDetails = CartItemDetailToCanopus(cart.ItemDetails)

	result.CartDetails.ID = cart.CartDetail.ID
	result.CartDetails.Payment.Key = cart.CartDetail.Payment.Key
	result.CartDetails.Payment.Type = cart.CartDetail.Payment.Type
	result.CartDetails.Amount = cart.CartDetail.Amount
	result.CartDetails.Title = cart.CartDetail.Title
	result.CartDetails.Currency = cart.CartDetail.Currency
	result.CartDetails.ExpiredAt = cart.CartDetail.ExpiredAt

	result.CustomerDetails.FirstName = cart.CustomerDetail.FirstName
	result.CustomerDetails.LastName = cart.CustomerDetail.LastName
	result.CustomerDetails.Email = cart.CustomerDetail.Email
	result.CustomerDetails.Phone = convertPhoneNumber(cart.CustomerDetail.Phone)

	result.CustomerDetails.BillingAddress.Phone = convertPhoneNumber(cart.CustomerDetail.BillingAddress.Phone)
	result.CustomerDetails.BillingAddress.FirstName = cart.CustomerDetail.BillingAddress.FirstName
	result.CustomerDetails.BillingAddress.LastName = cart.CustomerDetail.BillingAddress.LastName
	result.CustomerDetails.BillingAddress.Address = cart.CustomerDetail.BillingAddress.Address
	result.CustomerDetails.BillingAddress.City = cart.CustomerDetail.BillingAddress.City
	result.CustomerDetails.BillingAddress.PostalCode = cart.CustomerDetail.BillingAddress.PostalCode

	result.CustomerDetails.ShippingAddress.FirstName = cart.CustomerDetail.ShippingAddress.FirstName
	result.CustomerDetails.ShippingAddress.LastName = cart.CustomerDetail.ShippingAddress.LastName
	result.CustomerDetails.ShippingAddress.Phone = cart.CustomerDetail.ShippingAddress.Phone
	result.CustomerDetails.ShippingAddress.Address = cart.CustomerDetail.ShippingAddress.Address
	result.CustomerDetails.ShippingAddress.City = cart.CustomerDetail.ShippingAddress.City
	result.CustomerDetails.ShippingAddress.PostalCode = cart.CustomerDetail.ShippingAddress.PostalCode

	result.URL.ReturnURL = cart.Callback.Return
	result.URL.CancelURL = cart.Callback.Cancel
	result.URL.NotificationURL = cart.Callback.Success

	result.ExtendInfo.AdditionalPrefix = cart.ExtendInfo.AdditionalPrefix

	return result
}

func CartItemDetailToCanopus(items []CartPayloadItemDetail) []canopusgo.CartPayloadItemDetail {
	var result []canopusgo.CartPayloadItemDetail

	for _, item := range items {
		var tmp canopusgo.CartPayloadItemDetail
		tmp.Name = item.Name
		tmp.Desc = item.Desc
		tmp.Price = item.Price
		tmp.Quantity = item.Quantity
		tmp.SKU = item.SKU
		tmp.AdditionalInfo.NoHandphone = item.AdditionalInfo.NoHandphone
		result = append(result, tmp)
	}

	return result
}

func CartResponseFromCanopus(resp canopusgo.CartResponse) CartResponse {
	var result CartResponse

	result.CartID = resp.CartID
	result.Link = resp.PayTo
	result.Amount = resp.Amount
	result.Bank = resp.Bank
	result.Number = resp.Number

	return result
}

func convertPhoneNumber(phoneNumb string) string {
	if len(phoneNumb) == 0 {
		return phoneNumb
	}

	if string(phoneNumb[0]) != "0" {
		return "0" + phoneNumb
	}

	return phoneNumb
}
