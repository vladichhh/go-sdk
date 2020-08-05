package conn

import (
	"testing"

	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/test"
	"github.com/LimePay/go-sdk/test/helper"
	"github.com/LimePay/go-sdk/types"
	gock "gopkg.in/h2non/gock.v1"
)

var pingMock = types.Ping{
	Status: "ok",
}

func TestNewClient(t *testing.T) {
	et := helper.WrapTesting(t)

	requester := http.NewRequester(test.Env, test.APIKey, test.APISecret)

	connectionClient := NewClient(requester)

	et.Assert(connectionClient.requester == requester, "Requester address does not match after init")
}

func TestPing(t *testing.T) {
	et := helper.WrapTesting(t)

	defer gock.Off()

	gock.New(test.Env).
		Get("/ping").
		Reply(200).
		JSON(pingMock)

	connectionClient := NewClient(http.NewRequester(test.Env, test.APIKey, test.APISecret))

	res, err := connectionClient.Test()

	et.Assert(res.Status == pingMock.Status, "Ping status does not match")
	et.Assert(err == nil, "Not expected error returned")
}
