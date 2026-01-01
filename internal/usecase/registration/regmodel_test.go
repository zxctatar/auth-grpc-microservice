package registration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegInput_Success(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	ri, err := NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.NoError(t, err)
	assert.Equal(t, ri.FirstName, firstName)
	assert.Equal(t, ri.MiddleName, middleName)
	assert.Equal(t, ri.LastName, lastName)
	assert.Equal(t, ri.Password, pass)
	assert.Equal(t, ri.Email, email)
}

func TestNewRegInput_InvalidFirstName(t *testing.T) {
	firstName := ""
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	_, err := NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.ErrorIs(t, err, ErrEmptyFirstName)
}

func TestNewRegInput_InvalidLastName(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := ""
	pass := "somePass"
	email := "mail@mail.ru"

	_, err := NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.ErrorIs(t, err, ErrEmptyLastName)
}

func TestNewRegInput_InvalidPasswordName(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := ""
	email := "mail@mail.ru"

	_, err := NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.ErrorIs(t, err, ErrEmptyPassword)
}

func TestNewRegInput_InvalidEmailName(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := ""

	_, err := NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.ErrorIs(t, err, ErrEmptyEmail)
}
