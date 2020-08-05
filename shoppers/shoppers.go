package shoppers

import (
	"fmt"

	"github.com/LimePay/go-sdk/consts"
	"github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/types"
)

type ShoppersClient interface {
	Create(shopperData types.Shopper) (*types.Shopper, error)
	Get(shopperID string) (*types.Shopper, error)
	GetAll() (*[]types.Shopper, error)
	Update(shopperID string, shopperData types.Shopper) (*types.Shopper, error)
	GetWalletToken(shopperID string) (*types.WalletToken, error)
}

// BaseShoppersClient provides functionality to consume Shoppers resource
type BaseShoppersClient struct {
	requester http.Requester
}

// NewClient -
func NewClient(requester http.Requester) *BaseShoppersClient {
	return &BaseShoppersClient{requester}
}

// Create registeres a new shopper
func (s *BaseShoppersClient) Create(shopperData types.Shopper) (*types.Shopper, error) {
	shopper := &types.Shopper{}

	err := s.checkShopperVendor(&shopperData)

	if err != nil {
		return shopper, err
	}

	err = s.ExecuteRequest(consts.HTTPPost, consts.RouteCreateShopper, shopperData, shopper)

	return shopper, err
}

func (s *BaseShoppersClient) checkShopperVendor(shopperData *types.Shopper) error {
	if shopperData.Vendor == "" {
		vendors, err := s.getAllVendors()

		if err != nil {
			return err
		}

		if len(vendors) < 1 {
			return errors.NoVendorError
		}

		shopperData.Vendor = vendors[0].ID
	}

	return nil
}

// GetAllVendors retrieves all registered vendors
func (s *BaseShoppersClient) getAllVendors() ([]types.Vendor, error) {
	vendors := []types.Vendor{}

	err := s.ExecuteRequest(consts.HTTPGet, consts.RouteGetAllVendors, nil, &vendors)

	return vendors, err
}

// Get retrieves details for a given shopper
func (s *BaseShoppersClient) Get(shopperID string) (*types.Shopper, error) {
	route := fmt.Sprintf(consts.RouteGetShopper, shopperID)

	shopper := &types.Shopper{}

	err := s.ExecuteRequest(consts.HTTPGet, route, nil, shopper)

	return shopper, err
}

// GetAll retrieves all registered shoppers
func (s *BaseShoppersClient) GetAll() (*[]types.Shopper, error) {
	shoppers := &[]types.Shopper{}

	err := s.ExecuteRequest(consts.HTTPGet, consts.RouteGetAllShoppers, nil, shoppers)

	return shoppers, err
}

// Update updates details of a given shopper
func (s *BaseShoppersClient) Update(shopperID string, shopperData types.Shopper) (*types.Shopper, error) {
	route := fmt.Sprintf(consts.RoutePatchShopper, shopperID)

	shopper := &types.Shopper{}

	err := s.ExecuteRequest(consts.HTTPPatch, route, shopperData, shopper)

	return shopper, err
}

// GetWalletToken retrieves generated JSON wallet token for a given shopper
func (s *BaseShoppersClient) GetWalletToken(shopperID string) (*types.WalletToken, error) {
	route := fmt.Sprintf(consts.RouteGetWalletToken, shopperID)

	walletToken := &types.WalletToken{}

	err := s.ExecuteRequest(consts.HTTPGet, route, nil, walletToken)

	return walletToken, err
}

// ExecuteRequest -
func (s *BaseShoppersClient) ExecuteRequest(method string, route string, data interface{}, model interface{}) error {
	return s.requester.ExecuteRequest(method, route, data, model)
}
