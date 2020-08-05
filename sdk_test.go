package limepaysdk

import (
	"testing"

	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"
)

var pingMock = types.Ping{
	Status: "ok",
}

func TestConnect(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/ping").
		Reply(200).
		JSON(pingMock)

	limePaySDK, err := Connect(test.Env, test.APIKey, test.APISecret)

	et.Assert(limePaySDK != nil, "SDK instance could not be created")
	et.Assert(limePaySDK.Shoppers != nil, "ShoppersClient has not been instantiated")
	et.Assert(limePaySDK.Payments != nil, "PaymentsClient has not been instantiated")
	et.Assert(limePaySDK.FiatPayments != nil, "FiatPaymentsClient has not been instantiated")
	et.Assert(limePaySDK.RelayedPayments != nil, "RelayedPaymentsClient has not been instantiated")
	et.Assert(err == nil, "Not expected error returned")
}

func TestConnectThrowingError(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/ping").
		Reply(401).
		JSON("Unauthorized request")

	limePaySDK, err := Connect(test.Env, test.APIKey, test.APISecret)

	et.Assert(limePaySDK == nil, "SDK instance should not be created")
	et.Assert(err != nil, "Expected error to be returned")
}
