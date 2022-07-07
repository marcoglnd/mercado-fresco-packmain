package domain

import (
	"errors"
)

var (
	ErrIDNotFound = errors.New("product id not found")
	ErrEmptyID = errors.New("strconv.ParseInt: parsing \"\": invalid syntax")
)
