package jwtservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJWTService_Success(t *testing.T) {
	secretKey := []byte("key")
	timeOut := 5 * time.Second

	jwtService := NewJWTService(secretKey, &timeOut)

	assert.Equal(t, jwtService.SecretKey, secretKey)
	assert.Equal(t, *jwtService.TimeOut, timeOut)
}

func TestGenerateAndValidateToken_Success(t *testing.T) {
	secretKey := []byte("key")
	timeOut := 5 * time.Second
	userId := uint32(1)

	jwtService := NewJWTService(secretKey, &timeOut)

	token, err := jwtService.Generate(userId)

	assert.NoError(t, err)

	id, err := jwtService.ValidateToken(token)

	assert.NoError(t, err)
	assert.Equal(t, userId, id)
}

func TestGenerateAndValidateToken_FailValidateToken(t *testing.T) {
	secretKey := []byte("key")
	timeOut := 5 * time.Second
	userId := uint32(1)

	jwtService := NewJWTService(secretKey, &timeOut)

	_, err := jwtService.Generate(userId)

	assert.NoError(t, err)

	invalidToken := "asd"

	_, err = jwtService.ValidateToken(invalidToken)

	assert.Error(t, err)
}

func TestGenerateAndValidateToken_TimeOut(t *testing.T) {
	secretKey := []byte("key")
	timeOut := 1 * time.Millisecond
	userId := uint32(1)

	jwtService := NewJWTService(secretKey, &timeOut)

	token, err := jwtService.Generate(userId)

	assert.NoError(t, err)

	time.Sleep(2 * time.Millisecond)

	_, err = jwtService.ValidateToken(token)

	assert.Error(t, err)
}
