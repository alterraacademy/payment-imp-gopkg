package paymentgo

import "time"

// InitService ...
type CanopusService struct {
	CanopusType string
	MerchantKey []byte
	MerchantPem []byte
	MerchantID  string
	Secret      string
	TimeOut     time.Duration
}

// PaymentMethod ...
type PaymentMethod struct {
	From        string
	Key         string
	Name        string
	Type        string
	Logo        string
	Instruction interface{}
}

// CartPayload ...
type CartPayload struct {
	CartDetail struct {
		ID      string
		Payment struct {
			Key  string
			Type string
		}
		Amount    float64
		Title     string
		Currency  string
		ExpiredAt string
	}
	ItemDetails    []CartPayloadItemDetail
	CustomerDetail struct {
		FirstName      string
		LastName       string
		Email          string
		Phone          string
		BillingAddress struct {
			FirstName  string
			LastName   string
			Phone      string
			Address    string
			City       string
			PostalCode string
		}
		ShippingAddress struct {
			FirstName  string
			LastName   string
			Phone      string
			Address    string
			City       string
			PostalCode string
		}
	}
	Environment struct {
		Agent   string
		Mode    string
		Os      string
		Version string
	}
	Callback struct {
		Return  string
		Cancel  string
		Success string
	}
	ExtendInfo struct {
		AdditionalPrefix string
	}
}

// CartPayloadItemDetail item cart detail
type CartPayloadItemDetail struct {
	Name           string
	Desc           string
	Price          float64
	Quantity       int
	SKU            string
	AdditionalInfo struct {
		NoHandphone string
	}
}

// CartResponse ...
type CartResponse struct {
	CartID string
	Link   string
	Amount string
	Bank   string
	Number string
}
