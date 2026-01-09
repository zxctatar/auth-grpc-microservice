package regerror

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)
