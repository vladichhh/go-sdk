package fiat

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

var fiatPaymentMock = types.Payment{
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
	Type: "FIAT_PAYMENT",
}

var fiatPaymentWithoutWeiAmountMock = types.Payment{
	Currency:            "USD",
	Shopper:             "0",
	Vendor:              "vendor123",
	Items:               []types.Item{},
	FundTxData:          types.FundTxData{},
	GenericTransactions: []types.GenericTransaction{},
	Type:                "FIAT_PAYMENT",
}

var invoiceMock = "sample invoice content"

var receiptMock = "sample receipt content"

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
		Post("/payments").
		Reply(200).
		JSON(fiatPaymentMock)

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := fiatPaymentsClient.Create(fiatPaymentMock, privateKeyMock)

	et.Assert(fiatPaymentMock.ID == res.ID, "Fiat payment ID does not match")
	et.Assert(fiatPaymentMock.Status == res.Status, "Fiat payment status does not match")
	et.Assert(fiatPaymentMock.Date == res.Date, "Fiat payment date does not match")
	et.Assert(fiatPaymentMock.Currency == res.Currency, "Fiat payment currency does not match")
	et.Assert(fiatPaymentMock.Shopper == res.Shopper, "Fiat payment shopper does not match")
	et.Assert(fiatPaymentMock.Vendor == res.Vendor, "Fiat payment vendor does not match")
	et.Assert(fiatPaymentMock.FundTxData.TokenAmount == res.FundTxData.TokenAmount, "Fiat payment tokenAmount does not match")
	et.Assert(fiatPaymentMock.FundTxData.WeiAmount == res.FundTxData.WeiAmount, "Fiat payment weiAmount does not match")
	et.Assert(fiatPaymentMock.Type == res.Type, "Fiat payment type does not match")
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

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	_, err := fiatPaymentsClient.Create(fiatPaymentWithoutWeiAmountMock, privateKeyMock)

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*errors.SDKError)

	et.Assert(sdkError.ErrName == errors.InvalidTokenAndWeiAmountProvided.ErrName, "No matching error name")
	et.Assert(sdkError.ErrCode == errors.InvalidTokenAndWeiAmountProvided.ErrCode, "No matching error code")
	et.Assert(sdkError.ErrMessage == errors.InvalidTokenAndWeiAmountProvided.ErrMessage, "No matching error message")
}

func TestCreateWithInvalidPrivateKey(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/metadata").
		MatchParam("shopperId", "0").
		Reply(200).
		JSON(shopperMetadataMock)

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	_, err := fiatPaymentsClient.Create(fiatPaymentMock, "0x123")

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*errors.SDKError)

	et.Assert(sdkError.ErrName == errors.SigningError.ErrName, "No matching error name")
	et.Assert(sdkError.ErrCode == errors.SigningError.ErrCode, "No matching error code")
	et.Assert(sdkError.ErrMessage == errors.SigningError.ErrMessage, "No matching error message")
}

func TestGetInvoice(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/0/invoice/preview").
		Reply(200).
		JSON(invoiceMock)

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := fiatPaymentsClient.GetInvoice(fiatPaymentMock.ID)

	et.Assert(invoiceMock == res, "Invoice content does not match")
	et.Assert(err == nil, "Not expected error expected")
}

func TestSendInvoice(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/0/invoice").
		Reply(200)

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	err := fiatPaymentsClient.SendInvoice(fiatPaymentMock.ID)

	et.Assert(err == nil, "Not expected error expected")
}

func TestGetReceipt(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/0/receipt").
		Reply(200).
		JSON(receiptMock)

	fiatPaymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := fiatPaymentsClient.GetReceipt(fiatPaymentMock.ID)

	et.Assert(receiptMock == res, "Receipt content does not match")
	et.Assert(err == nil, "Not expected error expected")
}
