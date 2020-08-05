package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/LimePay/go-sdk/errors"
)

// Requester -
type Requester interface {
	ExecuteRequest(method string, route string, data interface{}, model interface{}) error
}

// BaseRequester -
type BaseRequester struct {
	env       string
	apiKey    string
	apiSecret string
}

// NewRequester -
func NewRequester(env string, apiKey string, apiSecret string) *BaseRequester {
	return &BaseRequester{env, apiKey, apiSecret}
}

// ExecuteRequest -
func (r *BaseRequester) ExecuteRequest(method string, route string, data interface{}, model interface{}) error {
	client := &http.Client{}

	buf := new(bytes.Buffer)

	if data != nil {
		json.NewEncoder(buf).Encode(&data)
	}

	req, err := http.NewRequest(method, join(r.env, route), buf)

	if err != nil {
		return err
	}

	req.SetBasicAuth(r.apiKey, r.apiSecret)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		sdkErr := errors.SDKError{}
		json.Unmarshal(bytes, &sdkErr)
		return &sdkErr
	}

	if model == nil {
		return nil
	}

	switch reflect.TypeOf(model).String() {
	case "*string":
		*(model.(*string)) = string(bytes)
	default:
		err = json.Unmarshal(bytes, model)

		if err != nil {
			return err
		}
	}

	return nil
}

func join(strs ...string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(str)
	}

	return sb.String()
}
