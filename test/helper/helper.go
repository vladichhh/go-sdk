package helper

import (
	"testing"
)

// ExtendedTesting -
type ExtendedTesting struct {
	*testing.T
}

// Assert -
func (et *ExtendedTesting) Assert(condition bool, msg ...interface{}) {

	if !condition {
		et.Error(msg...)
	}
}

// WrapTesting -
func WrapTesting(t *testing.T) *ExtendedTesting {
	return &ExtendedTesting{t}
}
