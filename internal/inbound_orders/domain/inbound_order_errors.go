package domain

import "errors"

var (
	ErrIdNotFound     = errors.New("employee id not found")
	ErrInvalidService = errors.New("invalid service")
)
