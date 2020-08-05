package limepaysdk

import (
	"github.com/LimePay/go-sdk/conn"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/payments"
	"github.com/LimePay/go-sdk/payments/fiat"
	"github.com/LimePay/go-sdk/payments/relayed"
	"github.com/LimePay/go-sdk/shoppers"
)

// LimePaySDK -
type LimePaySDK struct {
	Shoppers        shoppers.ShoppersClient
	Payments        payments.PaymentsClient
	FiatPayments    fiat.FiatPaymentsClient
	RelayedPayments relayed.RelayedPaymentsClient
}

// Connect connecting to LimePay API
func Connect(env string, apiKey string, apiSecret string) (*LimePaySDK, error) {
	requester := http.NewRequester(env, apiKey, apiSecret)

	err := checkConnection(requester)

	if err != nil {
		return nil, err
	}

	shoppers := shoppers.NewClient(requester)
	payments := payments.NewClient(requester)
	fiatPayments := fiat.NewClient(requester)
	relayedPayments := relayed.NewClient(requester)

	return &LimePaySDK{shoppers, payments, fiatPayments, relayedPayments}, nil
}

func checkConnection(requester http.Requester) error {
	connection := conn.NewClient(requester)

	_, err := connection.Test()

	return err
}
