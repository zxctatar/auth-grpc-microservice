package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewPostgresModel_Success(t *testing.T) {
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

func TestNewPostgresUserAuthDataModel_Su(t *testing.T) {
	id := uint32(1)
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)

	posModel := NewPostgresUserAuthDataModel(id, string(hashPassword))

	assert.Equal(t, id, posModel.Id)
	assert.Equal(t, string(hashPassword), posModel.HashPassword)
}
