package assert

import "testing"

//Assert wraps a testing.T
type Assert struct {
	t *testing.T
}

// NewAssert returns an Assert type.
func NewAssert(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) Equal(actual interface{}, expected interface{}, msg string) {
	if actual != expected {
		a.t.Errorf("%s not equal, actual: '%v', expected: '%v'", msg, actual, expected)
	}
}
