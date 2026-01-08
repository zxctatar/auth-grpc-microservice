package jwtservice

import "errors"

var (
	ErrInvalidToken = errors.New("token is invalid")
)
