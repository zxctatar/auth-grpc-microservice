package jwtservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJWTService_Success(t *testing.T) {
	secretKet := "key"
	timeOut := 5 * time.Second

	jwtService := NewJWTService(secretKet, &timeOut)

	assert.Equal(t, jwtService.SecretKey, secretKet)
	assert.Equal(t, *jwtService.TimeOut, timeOut)
}

func TestGenerate_Success(t *testing.T) {
	secretKet := "key"
	timeOut := 5 * time.Second

	jwtService := NewJWTService(secretKet, &timeOut)

	_, err := jwtService.Generate(1)

	assert.NoError(t, err)
}