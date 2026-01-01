package usecase

import "errors"

var (
	ErrEmptyFirstName = errors.New("empty first name")
	ErrEmptyLastName  = errors.New("empty last name")
	ErrEmptyPassword  = errors.New("empty password")
	ErrEmptyEmail     = errors.New("empty email")
)
