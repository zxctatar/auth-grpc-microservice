package userdomain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUserDomain_ValidInput(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)
	email := "mail@mail.ru"

	ud, err := NewUserDomain(
		firstName,
		middleName,
		lastName,
		string(hashPass),
		email,
	)

	assert.NoError(t, err)
	assert.Equal(t, ud.FirstName, firstName)
	assert.Equal(t, ud.MiddleName, middleName)
	assert.Equal(t, ud.LastName, lastName)
	assert.Equal(t, ud.HashPassword, string(hashPass))
	assert.Equal(t, ud.Email, email)
}

func TestNewUserDomain_InvalidEmail(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)
	email := "@mail.ru"

	_, err := NewUserDomain(
		firstName,
		middleName,
		lastName,
		string(hashPass),
		email,
	)

	assert.Error(t, err, ErrInvalidEmail)

	email = "mail@"

	_, err = NewUserDomain(
		firstName,
		middleName,
		lastName,
		string(hashPass),
		email,
	)

	assert.ErrorIs(t, err, ErrInvalidEmail)

	email = "mail@.ru"

	_, err = NewUserDomain(
		firstName,
		middleName,
		lastName,
		string(hashPass),
		email,
	)

	assert.ErrorIs(t, err, ErrInvalidEmail)
}

func TestRestoreUserDomain(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)
	email := "mail@mail.ru"

	ud := RestoreUserDomain(
		firstName,
		middleName,
		lastName,
		string(hashPass),
		email,
	)

	assert.Equal(t, ud.FirstName, firstName)
	assert.Equal(t, ud.MiddleName, middleName)
	assert.Equal(t, ud.LastName, lastName)
	assert.Equal(t, ud.HashPassword, string(hashPass))
	assert.Equal(t, ud.Email, email)
}