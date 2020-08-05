package errors

// SDKError -
type SDKError struct {
	ErrName    string `json:"errorName"`
	ErrCode    int    `json:"code"`
	ErrMessage string `json:"message"`
}

func (e *SDKError) Error() string {
	return e.ErrMessage
}

// SigningError -
var SigningError = &SDKError{"SIGNING_ERROR", 1011, "Could not sign authorization signature. Invalid parameters provided."}

// InvalidTokenAndWeiAmountProvided -
var InvalidTokenAndWeiAmountProvided = &SDKError{"VALIDATION_ERROR", 1017, "Invalid fundTxData object provided. tokenAmount or weiAmount cannot be undefined"}

// InvalidWeiAmountProvided -
var InvalidWeiAmountProvided = &SDKError{"VALIDATION_ERROR", 1018, "Invalid fundTxData object provided. weiAmount cannot be undefined"}

// NoVendorError -
var NoVendorError = &SDKError{"NO_VENDOR_ERROR", 1019, "You are required to have vendor in order to perform this operation"}
