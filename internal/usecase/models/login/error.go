package logmodel

import "errors"

var (
	ErrEmptyEmail = errors.New("empty email")
	ErrEmptyPassword = errors.New("empty password")
)