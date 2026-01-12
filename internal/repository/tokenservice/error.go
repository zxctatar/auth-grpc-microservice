package tokenservice

import "errors"

var (
	ErrInvalidSignature = errors.New("invalid token signature")
	ErrTokenMalformed   = errors.New("token malformed")
	ErrInvalidToken     = errors.New("token is invalid")
)
