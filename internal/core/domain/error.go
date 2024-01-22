package domain

import (
	"errors"
)

var (
	ErrBadRequest   = errors.New("bad request")
	ErrDataNotFound = errors.New("data not found")
)
