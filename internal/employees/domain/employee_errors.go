package domain

import "errors"

var (
	ErrIdNotFound               = errors.New("employee id not found")
	ErrDuplicatedID             = errors.New("duplicated card_number_id")
	ErrInvalidService           = errors.New("invalid service")
	CardNumberIdIsRequired      = errors.New("card_number_id is required")
	FirstNameIsRequired         = errors.New("first_name is required")
	LastNameIsRequired          = errors.New("last_name is required")
	WarehouseIdCannotBeNegative = errors.New("warehouse_id cannot be negative")
)
