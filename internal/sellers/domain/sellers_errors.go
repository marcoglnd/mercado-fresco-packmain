package domain

import "errors"

var (
	ErrIDNotFound    = errors.New("seller id not found")
	ErrDuplicatedCID = errors.New("duplicated cid")
)
