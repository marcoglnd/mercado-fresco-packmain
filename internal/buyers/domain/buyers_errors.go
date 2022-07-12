package domain

import "errors"

var (
	ErrIDNotFound   = errors.New("buyer id not found")
	ErrDuplicatedID = errors.New("duplicated card_number_id")
)
