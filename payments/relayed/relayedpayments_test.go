package relayed

import (
	"testing"

	"github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"
)

var shopperMetadataMock = types.SignatureMetadata{
	Nonce:          "0",
	ShopperAddress: "0x123",
	EscrowAddress:  "0x321",
}

var relayedPaymentMock = types.Payment{
	ID:       "0",
	Status:   "NEW",
	Date:     "01-01-2019",
	Currency: "USD",
	Shopper:  "0",
	Vendor:   "vendor123",
	Items: []types.Item{
		{
			Description: "Some description",
			LineAmount:  100.4,
			Quantity:    1,
		},
	},
	FundTxData: types.FundTxData{
		TokenAmount: "10000000000000000000",
		WeiAmount:   "60000000000000000",
	},
	GenericTransactions: []types.GenericTransaction{
		{
			GasPrice:       "1",
			GasLimit:       4700000,
			To:             "0x37688cFc875DC6AA6D39fE8449A759e434a86482",
			FunctionName:   "buySomeService",
			FunctionParams: []types.FunctionParams{},
		},
	},
	Type: "RELAYED_PAYMENT",
}

var relayedPaymentWithoutWeiAmountMock = types.Payment{
	Currency:            "USD",
	Shopper:             "0",
	Vendor:              "vendor123",
	Items:               []types.Item{},
	FundTxData:          types.FundTxData{},
	GenericTransactions: []types.GenericTransaction{},
	Type:                "RELAYED_PAYMENT",
}

var privateKeyMock = "d723d3cdf932464de15845c0719ca13ce15e64c83625d86ddbfc217bd2ac5f5a"

func TestCreate(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/metadata").
		MatchParam("shopperId", "0").
		Reply(200).
		JSON(shopperMetadataMock)

	gock.New(test.Env).
		Post("/payments/relayed").
		Reply(200).
		JSON(relayedPaymentMock)

	relayedPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := relayedPaymentsClient.Create(relayedPaymentMock, privateKeyMock)

	et.Assert(relayedPaymentMock.ID == res.ID, "Relayed payment ID does not match")
	et.Assert(relayedPaymentMock.Status == res.Status, "Relayed payment status does not match")
	et.Assert(relayedPaymentMock.Date == res.Date, "Relayed payment date does not match")
	et.Assert(relayedPaymentMock.Currency == res.Currency, "Relayed payment currency does not match")
	et.Assert(relayedPaymentMock.Shopper == res.Shopper, "Relayed payment shopper does not match")
	et.Assert(relayedPaymentMock.Vendor == res.Vendor, "Relayed payment vendor does not match")
	et.Assert(relayedPaymentMock.FundTxData.TokenAmount == res.FundTxData.TokenAmount, "Relayed payment tokenAmount does not match")
	et.Assert(relayedPaymentMock.FundTxData.WeiAmount == res.FundTxData.WeiAmount, "Relayed payment weiAmount does not match")
	et.Assert(relayedPaymentMock.Type == res.Type, "Relayed payment type does not match")
	et.Assert(err == nil, "Not expected error expected")
}

func TestCreateWithoutWeiAmount(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/metadata").
		MatchParam("shopperId", "0").
		Reply(200).
		JSON(shopperMetadataMock)

	relayedPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	_, err := relayedPaymentsClient.Create(relayedPaymentWithoutWeiAmountMock, privateKeyMock)

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*errors.SDKError)

	et.Assert(sdkError.ErrName == errors.InvalidWeiAmountProvided.ErrName, "No matching error name")
	et.Assert(sdkError.ErrCode == errors.InvalidWeiAmountProvided.ErrCode, "No matching error code")
	et.Assert(sdkError.ErrMessage == errors.InvalidWeiAmountProvided.ErrMessage, "No matching error message")
}

func TestCreateWithInvalidPrivateKey(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/metadata").
		MatchParam("shopperId", "0").
		Reply(200).
		JSON(shopperMetadataMock)

	relayedPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	_, err := relayedPaymentsClient.Create(relayedPaymentMock, "0x123")

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*errors.SDKError)

	et.Assert(sdkError.ErrName == errors.SigningError.ErrName, "No matching error name")
	et.Assert(sdkError.ErrCode == errors.SigningError.ErrCode, "No matching error code")
	et.Assert(sdkError.ErrMessage == errors.SigningError.ErrMessage, "No matching error message")
}
