package http

import (
	"errors"
	"testing"

	"github.com/LimePay/go-sdk/consts"
	sdkerr "github.com/LimePay/go-sdk/errors"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"

	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
)

var authorizaionErrorMock = sdkerr.SDKError{
	ErrName:    "AUTHORIZATION_ERROR",
	ErrCode:    1111,
	ErrMessage: "Unauthorized request",
}

var internalServerErrorMock = sdkerr.SDKError{
	ErrName:    "INTERNAL_SERVER_ERROR",
	ErrCode:    5001,
	ErrMessage: "Something went wrong. Contact your support for more information",
}

func TestNewRequester(t *testing.T) {
	et := helper.WrapTesting(t)

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	et.Assert(requester != nil, "Requester has not been instantiated")
	et.Assert(test.Env == requester.env, "Requester env does not match")
	et.Assert(test.APIKey == requester.apiKey, "Requester apiKey does not match")
	et.Assert(test.APISecret == requester.apiSecret, "Requester apiSecret does not match")
}

func TestExecuteRequest(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/0").
		Reply(200).
		JSON(types.Shopper{})

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPGet, "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err == nil, "Not expected error returned")
}

func TestExecuteRequestWithData(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Patch("/shoppers/0").
		Reply(200).
		JSON(types.Shopper{})

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPPatch, "/shoppers/0", &types.Shopper{}, &types.Shopper{})

	et.Assert(err == nil, "Not expected error returned")
}

func TestExecuteRequestWithWrongMethod(t *testing.T) {
	et := helper.WrapTesting(t)

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest("?", "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err != nil, "Expected error to be returned")
}

func TestExecuteRequestWithProtocolError(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/0").
		ReplyError(errors.New("Protocol error"))

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPGet, "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err != nil, "Expected error to be returned")
}

func TestExecuteRequestWithUnmarshallingErr(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/0").
		Reply(200).
		JSON([]types.Shopper{})

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPGet, "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err != nil, "Expected error to be returned")
}

func TestExecuteRequestThrowingAuthorizationErr(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/0").
		Reply(401).
		JSON(authorizaionErrorMock)

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPGet, "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*sdkerr.SDKError)

	et.Assert(authorizaionErrorMock.ErrName == sdkError.ErrName, "Error name does not match")
	et.Assert(authorizaionErrorMock.ErrCode == sdkError.ErrCode, "Error code does not match")
	et.Assert(authorizaionErrorMock.ErrMessage == sdkError.ErrMessage, "Error message does not match")
}

func TestExecuteRequestThrowingInternalServerErr(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/shoppers/0").
		Reply(500).
		JSON(internalServerErrorMock)

	requester := NewRequester(test.Env, test.APIKey, test.APISecret)

	err := requester.ExecuteRequest(consts.HTTPGet, "/shoppers/0", nil, &types.Shopper{})

	et.Assert(err != nil, "Expected error to be returned")

	sdkError := err.(*sdkerr.SDKError)

	et.Assert(internalServerErrorMock.ErrName == sdkError.ErrName, "Error name does not match")
	et.Assert(internalServerErrorMock.ErrCode == sdkError.ErrCode, "Error code does not match")
	et.Assert(internalServerErrorMock.ErrMessage == sdkError.ErrMessage, "Error message does not match")
}
