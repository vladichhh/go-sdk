package relayed

import (
	"github.com/LimePay/go-sdk/consts"
	"github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/payments"
	"github.com/LimePay/go-sdk/types"
)

var relayedSignatureValueTypes = []string{"uint256", "address", "address", "uint256"}

type internalPaymentsClient interface {
	payments.PaymentsClient

	CreatePayment(route string, paymentData types.Payment) (*types.Payment, error)
	GetSignatureMetadata(shopperID string) (types.SignatureMetadata, error)
	Sign(privateKey string, paramTypes []string, paramValues []string) (string, error)
	ExecuteRequest(method string, route string, data interface{}, model interface{}) error
}

// RelayedPaymentsClient -
type RelayedPaymentsClient interface {
	payments.RichPaymentsClient
}

// BaseRelayedPaymentsClient -
type BaseRelayedPaymentsClient struct {
	internalPaymentsClient
}

// NewClient -
func NewClient(requester http.Requester) *BaseRelayedPaymentsClient {
	return &BaseRelayedPaymentsClient{payments.NewClient(requester)}
}

// Create -
func (r *BaseRelayedPaymentsClient) Create(paymentData types.Payment, privateKey string) (*types.Payment, error) {
	payment := &types.Payment{}

	if paymentData.FundTxData.AuthorizationSignature == "" {
		signatureMetadata, err := r.GetSignatureMetadata(paymentData.Shopper)

		if err != nil {
			return payment, err
		}

		authorizationSignature, errr := r.computeAuthorizationSignature(signatureMetadata, paymentData.FundTxData, privateKey)

		if errr != nil {
			return payment, errr
		}

		paymentData.FundTxData.Nonce = signatureMetadata.Nonce
		paymentData.FundTxData.AuthorizationSignature = authorizationSignature
	}

	return r.CreatePayment(consts.RouteCreateRelayedPayment, paymentData)
}

func (r *BaseRelayedPaymentsClient) computeAuthorizationSignature(signatureMetadata types.SignatureMetadata, fundTxData types.FundTxData, privateKey string) (string, error) {
	if fundTxData.WeiAmount == "" {
		return "", errors.InvalidWeiAmountProvided
	}

	res, err := r.Sign(privateKey, relayedSignatureValueTypes,
		[]string{signatureMetadata.Nonce, signatureMetadata.EscrowAddress, signatureMetadata.ShopperAddress, "", fundTxData.WeiAmount})

	if err != nil {
		return "", errors.SigningError
	}

	return res, nil
}
