package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	// t.Helper() indicates to the go test runner that Equal() is
	// a test helper.
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; expected %v", actual, expected)
	}
}