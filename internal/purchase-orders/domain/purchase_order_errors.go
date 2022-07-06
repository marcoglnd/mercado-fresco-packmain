package domain

import "errors"

var (
	ErrDuplicatedID = errors.New("duplicated card_number_id")
)
