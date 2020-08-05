package shoppers

import (
	"testing"

	"github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"
)

var shopperMock = types.Shopper{
	ID:                "12345",
	Vendor:            "vendor123",
	FirstName:         "Alexander",
	LastName:          "Kostov",
	Email:             "s.kostov@abv.bg",
	WalletAddress:     "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc034",
	MaliciousAttempts: 0,
}

var vendorsMock = []types.Vendor{
	{
		ID:                    "0",
		Name:                  "Test",
		Email:                 "test@abv.bg",
		Country:               "Bulgaria",
		City:                  "Sofia",
		FirstName:             "Lime",
		LastName:              "Chain",
		Address:               "Pernik",
		Phone:                 "+359123123123",
		Zip:                   "1000",
		State:                 "Sofia",
		VatID:                 "abc123abc",
		TaxID:                 "123abc123",
		PayoutInfo:            []types.PayoutInfo{},
		VendorPrincipal:       types.VendorPrincipal{},
		ReceiptEmail:          "test@abv.bg",
		EmailSetup:            types.EmailSetup{},
		DefaultPayoutCurrency: "bgn",
		AutoReceipt:           true,
		AutoInvoice:           true,
		Frequency:             "daily",
		RawLogo:               "photo",
	},
	{
		ID:                    "1",
		Name:                  "Test1",
		Email:                 "test@abv.bg1",
		Country:               "Bulgaria1",
		City:                  "Sofia1",
		FirstName:             "Lime1",
		LastName:              "Chain1",
		Address:               "Pernik1",
		Phone:                 "+3591231231231",
		Zip:                   "10001",
		State:                 "Sofia1",
		VatID:                 "abc123abc1",
		TaxID:                 "123abc1231",
		PayoutInfo:            []types.PayoutInfo{},
		VendorPrincipal:       types.VendorPrincipal{},
		ReceiptEmail:          "test@abv.bg1",
		EmailSetup:            types.EmailSetup{},
		DefaultPayoutCurrency: "bgn1",
		AutoReceipt:           true,
		AutoInvoice:           true,
		Frequency:             "daily1",
		RawLogo:               "photo1",
	},
}

var shoppersMock = []types.Shopper{
	{
		ID:                "0",
		Vendor:            "vendor123",
		FirstName:         "George",
		LastName:          "Ivanov",
		Email:             "g.ivanov@abv.bg",
		WalletAddress:     "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc032",
		MaliciousAttempts: 0,
	},
	{
		ID:                "1",
		Vendor:            "vendor123",
		FirstName:         "Alex",
		LastName:          "Petrov",
		Email:             "a.petrov@abv.bg",
		WalletAddress:     "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc033",
		MaliciousAttempts: 0,
	},
}

var shopperWithLPWalletMock = types.Shopper{
	ID:                "3",
	Vendor:            "vendor123",
	FirstName:         "Sasho",
	LastName:          "Kostov",
	Email:             "s.kostov@abv.bg",
	WalletAddress:     "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc034",
	MaliciousAttempts: 0,
	UseLimePayWallet:  true,
}

var walletTokenMock = types.WalletToken{
	WalletToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlVc2VyIjoiNWMxOGRiYmE3ZGMxOGQzMGExZTk3OGYyIiwidXNlciI6IjVjMThkYjhkN2RjMThkMzBhMWU5NzhlYyIsInNob3BwZXJJZCI6IjVjNTJiYjQwMjRjNjk2NWY0N2NhN2ZhMiIsImNyZWF0ZWRPbiI6MTU1MDY1MjI4MjcwMywiaWF0IjoxNTUwNjUyMjgyLCJleHAiOjE1NTA2NTU4ODJ9.83Jx7K_FenuTcSQNrp8s5Xfl_g2NutD1RmSV23d5_Fk",
}

func TestNewClient(t *testing.T) {
	et := helper.WrapTesting(t)

	requester := http.NewRequester(test.Env, test.APIKey, test.APISecret)

	shoppersClient := NewClient(requester)

	et.Assert(shoppersClient.requester == requester, "Requester address does not match after init")
}

func TestCreateWithSpecifyingVendor(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Post("/shoppers").
		Reply(200).
		JSON(shopperMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.Create(types.Shopper{
		Vendor:        "vendor123",
		FirstName:     "Alexander",
		LastName:      "Kostov",
		Email:         "s.kostov@abv.bg",
		WalletAddress: "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc034"})

	et.Assert(shopperMock.ID == res.ID, "Shopper ID does not match")
	et.Assert(shopperMock.Vendor == res.Vendor, "Shopper vendor does not match")
	et.Assert(shopperMock.FirstName == res.FirstName, "Shopper firstName does not match")
	et.Assert(shopperMock.LastName == res.LastName, "Shopper lastName does not match")
	et.Assert(shopperMock.Email == res.Email, "Shopper email does not match")
	et.Assert(shopperMock.WalletAddress == res.WalletAddress, "Shopper walletAddress does not match")
	et.Assert(shopperMock.MaliciousAttempts == res.MaliciousAttempts, "Shopper maliciousAttempt does not match")
	et.Assert(err == nil, "Not expected error returned")
}

func TestCreateWithoutSpecifyingVendor(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/vendors").
		Reply(200).
		JSON(vendorsMock)

	gock.New(test.Env).
		Post("/shoppers").
		Reply(200).
		JSON(shopperMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.Create(types.Shopper{
		FirstName:     "Alexander",
		LastName:      "Kostov",
		Email:         "s.kostov@abv.bg",
		WalletAddress: "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc034"})

	et.Assert(shopperMock.ID == res.ID, "Shopper ID does not match")
	et.Assert(shopperMock.Vendor == res.Vendor, "Shopper vendor does not match")
	et.Assert(shopperMock.FirstName == res.FirstName, "Shopper firstName does not match")
	et.Assert(shopperMock.LastName == res.LastName, "Shopper lastName does not match")
	et.Assert(shopperMock.Email == res.Email, "Shopper email does not match")
	et.Assert(shopperMock.WalletAddress == res.WalletAddress, "Shopper walletAddress does not match")
	et.Assert(shopperMock.MaliciousAttempts == res.MaliciousAttempts, "Shopper maliciousAttempt does not match")
	et.Assert(err == nil, "Not expected error returned")
}

func TestCreateNoRegisteredVendors(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/vendors").
		Reply(200).
		JSON([]types.Vendor{})

	gock.New(test.Env).
		Post("/shoppers").
		Reply(200).
		JSON(shopperMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	_, err := shoppersClient.Create(types.Shopper{
		FirstName:     "Alexander",
		LastName:      "Kostov",
		Email:         "s.kostov@abv.bg",
		WalletAddress: "0x8E8FD30C784BBb9B80877052AAE4bd9D43BCc034"})

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*errors.SDKError)

	et.Assert(sdkError.ErrName == errors.NoVendorError.ErrName, "No matching error name")
	et.Assert(sdkError.ErrCode == errors.NoVendorError.ErrCode, "No matching error code")
	et.Assert(sdkError.ErrMessage == errors.NoVendorError.ErrMessage, "No matching error message")
}

func TestGet(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/12345").
		Reply(200).
		JSON(shopperMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.Get(shopperMock.ID)

	et.Assert(shopperMock.ID == res.ID, "Shopper ID does not match")
	et.Assert(shopperMock.Vendor == res.Vendor, "Shopper vendor does not match")
	et.Assert(shopperMock.FirstName == res.FirstName, "Shopper firstName does not match")
	et.Assert(shopperMock.LastName == res.LastName, "Shopper lastName does not match")
	et.Assert(shopperMock.Email == res.Email, "Shopper email does not match")
	et.Assert(shopperMock.WalletAddress == res.WalletAddress, "Shopper walletAddress does not match")
	et.Assert(shopperMock.MaliciousAttempts == res.MaliciousAttempts, "Shopper maliciousAttempt does not match")
	et.Assert(err == nil, "Not expected error returned")
}

func TestGetAll(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers").
		Reply(200).
		JSON(shoppersMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.GetAll()

	for i, s := range *res {
		et.Assert(shoppersMock[i].ID == s.ID, "Shopper ID does not match")
		et.Assert(shoppersMock[i].Vendor == s.Vendor, "Shopper vendor does not match")
		et.Assert(shoppersMock[i].FirstName == s.FirstName, "Shopper firstName does not match")
		et.Assert(shoppersMock[i].LastName == s.LastName, "Shopper lastName does not match")
		et.Assert(shoppersMock[i].Email == s.Email, "Shopper email does not match")
		et.Assert(shoppersMock[i].WalletAddress == s.WalletAddress, "Shopper walletAddress does not match")
		et.Assert(shoppersMock[i].MaliciousAttempts == s.MaliciousAttempts, "Shopper maliciousAttempt does not match")
	}

	et.Assert(err == nil, "Not expected error returned")
}

func TestUpdate(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Patch("/shoppers/12345").
		Reply(200).
		JSON(shopperMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.Update(shopperMock.ID, shopperMock)

	et.Assert(shopperMock.ID == res.ID, "Shopper ID does not match")
	et.Assert(shopperMock.Vendor == res.Vendor, "Shopper vendor does not match")
	et.Assert(shopperMock.FirstName == res.FirstName, "Shopper firstName does not match")
	et.Assert(shopperMock.LastName == res.LastName, "Shopper lastName does not match")
	et.Assert(shopperMock.Email == res.Email, "Shopper email does not match")
	et.Assert(shopperMock.WalletAddress == res.WalletAddress, "Shopper walletAddress does not match")
	et.Assert(shopperMock.MaliciousAttempts == res.MaliciousAttempts, "Shopper maliciousAttempt does not match")
	et.Assert(err == nil, "Not expected error returned")
}

func TestGetWalletToken(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/3/walletToken").
		Reply(200).
		JSON(walletTokenMock)

	shoppersClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := shoppersClient.GetWalletToken(shopperWithLPWalletMock.ID)

	et.Assert(walletTokenMock.WalletToken == res.WalletToken, "Wallet token does not match")
	et.Assert(err == nil, "Not expected error returned")
}
