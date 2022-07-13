package domain

import "errors"

var (
	ErrIDNotFound   = errors.New("section id not found")
	ErrDuplicatedID = errors.New("duplicated batch_number")
)
