package decoder

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := Error{
		Errors: map[string]error{
			"fieldname": errors.New("Failed to parse"),
		},
	}

	if want, got := "Error decoding: fieldname", err.Error(); want != got {
		t.Errorf("Error() should of been %q, but was %q", want, got)
	}
}
