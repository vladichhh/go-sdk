package payments

import (
	"fmt"
	"strings"

	"github.com/LimePay/go-sdk/consts"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PaymentsClient -
type PaymentsClient interface {
	Get(paymentID string) (*types.Payment, error)
	GetAll() (*[]types.Payment, error)
}

// RichPaymentsClient -
type RichPaymentsClient interface {
	PaymentsClient

	Create(paymentData types.Payment, privateKey string) (*types.Payment, error)
}

// BasePaymentsClient -
type BasePaymentsClient struct {
	Requester http.Requester
}

// NewClient -
func NewClient(requester http.Requester) *BasePaymentsClient {
	return &BasePaymentsClient{requester}
}

// Get -
func (p *BasePaymentsClient) Get(paymentID string) (*types.Payment, error) {
	route := fmt.Sprintf(consts.RouteGetPayment, paymentID)

	payment := &types.Payment{}

	err := p.ExecuteRequest(consts.HTTPGet, route, nil, payment)

	return payment, err
}

// GetAll -
func (p *BasePaymentsClient) GetAll() (*[]types.Payment, error) {
	payments := &[]types.Payment{}

	err := p.ExecuteRequest(consts.HTTPGet, consts.RouteGetAllPayments, nil, payments)

	return payments, err
}

// CreatePayment -
func (p *BasePaymentsClient) CreatePayment(route string, paymentData types.Payment) (*types.Payment, error) {
	payment := &types.Payment{}

	err := p.ExecuteRequest(consts.HTTPPost, route, paymentData, payment)

	return payment, err
}

// GetSignatureMetadata -
func (p *BasePaymentsClient) GetSignatureMetadata(shopperID string) (types.SignatureMetadata, error) {
	route := fmt.Sprintf(consts.RouteGetSignatureMetadata, shopperID)

	metadata := types.SignatureMetadata{}

	err := p.ExecuteRequest(consts.HTTPGet, route, nil, &metadata)

	return metadata, err
}

// Sign -
func (p *BasePaymentsClient) Sign(privateKey string, paramTypes []string, paramValues []string) (string, error) {

	pk, err := crypto.HexToECDSA(privateKey)

	if err != nil {
		return "", err
	}

	params := append(paramTypes, paramValues...)

	var sb strings.Builder

	for _, el := range params {
		sb.WriteString(el)
	}

	bytes := []byte(sb.String())

	hash := crypto.Keccak256Hash(bytes)

	signature, err := crypto.Sign(hash.Bytes(), pk)

	if err != nil {
		return "", err
	}

	return hexutil.Encode(signature), nil
}

// ExecuteRequest -
func (p *BasePaymentsClient) ExecuteRequest(method string, route string, data interface{}, model interface{}) error {
	return p.Requester.ExecuteRequest(method, route, data, model)
}
