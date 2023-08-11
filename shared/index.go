package shared

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidIndex = errors.New("shared: invalid index")
)

func IsValidIndex(idx string) error {
	idxInt, err := strconv.Atoi(idx)
	if err != nil {
		return ErrInvalidIndex
	}

	if idxInt < 1 || idxInt > 114 {
		return ErrInvalidIndex
	}

	return nil
}
