package loginerror

import "errors"

var (
	ErrWrongPassword = errors.New("wrong password")
)
