package payments

import (
	"testing"

	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"
)

var paymentMock = types.Payment{
	ID:       "0",
	Status:   "NEW",
	Date:     "01-01-2019",
	Currency: "USD",
	Shopper:  "5c810be19693cc2a02fa0dab",
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

var paymentsMock = []types.Payment{
	{
		ID:       "0",
		Status:   "NEW",
		Date:     "01-01-2019",
		Currency: "USD",
		Shopper:  "5c810be19693cc2a02fa0dab",
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
	},
	{
		ID:       "1",
		Status:   "NEW",
		Date:     "02-01-2019",
		Currency: "USD",
		Shopper:  "5c810be19693cc2a02fa0dab",
		Vendor:   "vendor123",
		Items: []types.Item{
			{
				Description: "Another description",
				LineAmount:  20.5,
				Quantity:    5,
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
	},
}

func TestNewClient(t *testing.T) {
	et := helper.WrapTesting(t)

	requester := http.NewRequester(test.Env, test.APIKey, test.APISecret)

	paymentsClient := NewClient(requester)

	et.Assert(paymentsClient.Requester == requester, "Requester address does not match after init")
}

func TestGet(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments/0").
		Reply(200).
		JSON(paymentMock)

	paymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	payment, err := paymentsClient.Get(paymentMock.ID)

	et.Assert(paymentMock.ID == payment.ID, "Payment ID does not match")
	et.Assert(paymentMock.Status == payment.Status, "Payment status does not match")
	et.Assert(paymentMock.Date == payment.Date, "Payment date does not match")
	et.Assert(paymentMock.Currency == payment.Currency, "Payment currency does not match")
	et.Assert(paymentMock.Shopper == payment.Shopper, "Payment shopper does not match")
	et.Assert(paymentMock.Vendor == payment.Vendor, "Payment vendor does not match")
	et.Assert(paymentMock.FundTxData.TokenAmount == payment.FundTxData.TokenAmount, "Payment tokenAmount does not match")
	et.Assert(paymentMock.FundTxData.WeiAmount == payment.FundTxData.WeiAmount, "Payment weiAmount does not match")
	et.Assert(paymentMock.Type == payment.Type, "Payment type does not match")
	et.Assert(err == nil, "Not expected error expected")
}

func TestGetAll(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/payments").
		Reply(200).
		JSON(paymentsMock)

	paymentsClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := paymentsClient.GetAll()

	for i, p := range *res {
		et.Assert(paymentsMock[i].ID == p.ID, "Payment ID does not match")
		et.Assert(paymentsMock[i].Status == p.Status, "Payment status does not match")
		et.Assert(paymentsMock[i].Date == p.Date, "Payment date does not match")
		et.Assert(paymentsMock[i].Currency == p.Currency, "Payment currency does not match")
		et.Assert(paymentsMock[i].Shopper == p.Shopper, "Payment shopper does not match")
		et.Assert(paymentsMock[i].Vendor == p.Vendor, "Payment vendor does not match")
		et.Assert(paymentsMock[i].FundTxData.TokenAmount == p.FundTxData.TokenAmount, "Payment tokenAmount does not match")
		et.Assert(paymentsMock[i].FundTxData.WeiAmount == p.FundTxData.WeiAmount, "Payment weiAmount does not match")
		et.Assert(paymentsMock[i].Type == p.Type, "Payment type does not match")
	}

	et.Assert(err == nil, "Not expected error returned")
}
