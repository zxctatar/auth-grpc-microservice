package logmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoginInput_Success(t *testing.T) {
	email := "mail@mail.ru"
	password := "somePass"

	li, err := NewLoginInput(email, password)

	assert.NoError(t, err)
	assert.Equal(t, li.Email, email)
	assert.Equal(t, li.Password, password)
}

func TestNewLoginInput_InvalidEmail(t *testing.T) {
	email := ""
	password := "somePass"

	_, err := NewLoginInput(email, password)

	assert.ErrorIs(t, err, ErrEmptyEmail)
}

func TestNewLoginInput_InvalidPassword(t *testing.T) {
	email := "mail@mail.ru"
	password := ""

	_, err := NewLoginInput(email, password)

	assert.ErrorIs(t, err, ErrEmptyPassword)
}