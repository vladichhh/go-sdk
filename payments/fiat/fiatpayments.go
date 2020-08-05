package fiat

import (
	"fmt"

	"github.com/LimePay/go-sdk/consts"
	"github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/payments"
	"github.com/LimePay/go-sdk/types"
)

var fiatSignatureValueTypes = []string{"uint256", "address", "address", "uint256", "uint256"}

type internalPaymentsClient interface {
	payments.PaymentsClient

	CreatePayment(route string, paymentData types.Payment) (*types.Payment, error)
	GetSignatureMetadata(shopperID string) (types.SignatureMetadata, error)
	Sign(privateKey string, paramTypes []string, paramValues []string) (string, error)
	ExecuteRequest(method string, route string, data interface{}, model interface{}) error
}

// FiatPaymentsClient -
type FiatPaymentsClient interface {
	payments.RichPaymentsClient

	GetInvoice(paymentID string) (string, error)
	SendInvoice(paymentID string) error
	GetReceipt(paymentID string) (string, error)
}

// BaseFiatPaymentsClient -
type BaseFiatPaymentsClient struct {
	internalPaymentsClient
}

// NewClient -
func NewClient(requester http.Requester) *BaseFiatPaymentsClient {
	return &BaseFiatPaymentsClient{payments.NewClient(requester)}
}

// Create -
func (f *BaseFiatPaymentsClient) Create(paymentData types.Payment, privateKey string) (*types.Payment, error) {
	payment := &types.Payment{}

	if paymentData.FundTxData.AuthorizationSignature == "" {
		signatureMetadata, err := f.GetSignatureMetadata(paymentData.Shopper)

		if err != nil {
			return payment, err
		}

		authorizationSignature, err := f.computeAuthorizationSignature(signatureMetadata, paymentData.FundTxData, privateKey)

		if err != nil {
			return payment, err
		}

		paymentData.FundTxData.Nonce = signatureMetadata.Nonce
		paymentData.FundTxData.AuthorizationSignature = authorizationSignature
	}

	return f.CreatePayment(consts.RouteCreateFiatPayment, paymentData)
}

func (f *BaseFiatPaymentsClient) computeAuthorizationSignature(signatureMetadata types.SignatureMetadata, fundTxData types.FundTxData, privateKey string) (string, error) {
	if fundTxData.TokenAmount == "" {
		fundTxData.TokenAmount = "0"
	}

	if fundTxData.WeiAmount == "" {
		return "", errors.InvalidTokenAndWeiAmountProvided
	}

	res, err := f.Sign(privateKey, fiatSignatureValueTypes,
		[]string{signatureMetadata.Nonce, signatureMetadata.EscrowAddress, signatureMetadata.ShopperAddress, fundTxData.TokenAmount, fundTxData.WeiAmount})

	if err != nil {
		return "", errors.SigningError
	}

	return res, nil
}

// GetInvoice -
func (f *BaseFiatPaymentsClient) GetInvoice(paymentID string) (string, error) {
	route := fmt.Sprintf(consts.RouteGetInvoice, paymentID)

	invoice := ""

	err := f.ExecuteRequest(consts.HTTPGet, route, nil, &invoice)

	return invoice, err
}

// SendInvoice -
func (f *BaseFiatPaymentsClient) SendInvoice(paymentID string) error {
	route := fmt.Sprintf(consts.RouteSendInvoice, paymentID)

	err := f.ExecuteRequest(consts.HTTPGet, route, nil, nil)

	return err
}

// GetReceipt -
func (f *BaseFiatPaymentsClient) GetReceipt(paymentID string) (string, error) {
	route := fmt.Sprintf(consts.RouteGetReceipt, paymentID)

	receipt := ""

	err := f.ExecuteRequest(consts.HTTPGet, route, nil, &receipt)

	return receipt, err
}
