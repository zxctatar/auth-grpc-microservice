package posmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresUserModel_Success(t *testing.T) {
	id := uint32(1)
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass := "somePass"
	email := "mail@mail.ru"

	posModel := NewPostgresModel(
		id,
		firstName,
		middleName,
		lastName,
		hashPass,
		email,
	)

	assert.Equal(t, posModel.Id, id)
	assert.Equal(t, posModel.FirstName, firstName)
	assert.Equal(t, posModel.MiddleName.String, middleName)
	assert.Equal(t, posModel.LastName, lastName)
	assert.Equal(t, posModel.HashPassword, hashPass)
	assert.Equal(t, posModel.Email, email)
}
