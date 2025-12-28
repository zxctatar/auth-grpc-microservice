package registration

import "errors"

var (
	// regmodel.go errors
	ErrEmptyFirstName = errors.New("empty first name")
	ErrEmptyLastName  = errors.New("empty last name")
	ErrEmptyPassword  = errors.New("empty password")
	ErrEmptyEmail     = errors.New("empty email")

	// register.go errors
	ErrUserAlreadyExists = errors.New("user already exists")
)
