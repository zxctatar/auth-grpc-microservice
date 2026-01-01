package postgres

import (
	userdomain "auth/internal/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelToDomain_Success(t *testing.T) {
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

	user := modelToDomain(posModel)

	assert.Equal(t, user.FirstName, posModel.FirstName)
	assert.Equal(t, user.MiddleName, posModel.MiddleName.String)
	assert.Equal(t, user.LastName, posModel.LastName)
	assert.Equal(t, user.HashPassword, posModel.HashPassword)
	assert.Equal(t, user.Email, posModel.Email)
}

func TestDomainToModel_Success(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass := "somePass"
	email := "mail@mail.ru"

	user := userdomain.RestoreUserDomain(
		firstName,
		middleName,
		lastName,
		hashPass,
		email,
	)

	posModel := domainToModel(user)

	assert.Equal(t, posModel.FirstName, user.FirstName)
	assert.Equal(t, posModel.MiddleName.String, user.MiddleName)
	assert.Equal(t, posModel.LastName, user.LastName)
	assert.Equal(t, posModel.HashPassword, user.HashPassword)
	assert.Equal(t, posModel.Email, user.Email)
}
