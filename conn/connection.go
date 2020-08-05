package conn

import (
	"github.com/LimePay/go-sdk/consts"
	"github.com/LimePay/go-sdk/http"
	"github.com/LimePay/go-sdk/types"
)

// ConnectionClient -
type ConnectionClient struct {
	requester http.Requester
}

// NewClient -
func NewClient(requester http.Requester) *ConnectionClient {
	return &ConnectionClient{requester}
}

// Test -
func (c *ConnectionClient) Test() (*types.Ping, error) {
	ping := &types.Ping{}

	err := c.requester.ExecuteRequest(consts.HTTPGet, consts.RoutePing, nil, ping)

	return ping, err
}
