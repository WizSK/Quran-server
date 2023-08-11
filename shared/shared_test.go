package shared

import (
	"errors"
	"testing"
)

func TestIsValidIndex(t *testing.T) {
	tests := []struct {
		idx string
		err error
	}{
		{
			idx: "foo",
			err: ErrInvalidIndex,
		},
		{
			idx: "",
			err: ErrInvalidIndex,
		},
		{
			idx: "0",
			err: ErrInvalidIndex,
		},
		{
			idx: "115",
			err: ErrInvalidIndex,
		},
		{
			idx: "1125",
			err: ErrInvalidIndex,
		},
	}

	for _, test := range tests {
		err := IsValidIndex(test.idx)
		if !errors.Is(err, test.err) {
			t.Errorf("expected err %q,got %q", test.err, err)
		}
	}
}
