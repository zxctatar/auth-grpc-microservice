package postgres

import (
	userdomain "auth/internal/domain/user"
	"auth/internal/repository"
	"context"
	"database/sql"
	"io"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSave_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	repo := NewPostgres(log, db)

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

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(firstName, middleName, lastName, hashPass, email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.Save(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, id, uint32(1))
}

func TestFindByEmail_UserFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	repo := NewPostgres(log, db)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	hashPass := "somePass"
	email := "mail@mail.ru"

	mock.ExpectQuery("SELECT id, first_name, middle_name, last_name, hash_password, email FROM users WHERE email = \\$1").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "middle_name", "last_name", "hash_password", "email"}).
			AddRow(1, firstName, middleName, lastName, hashPass, email))

	userOut, err := repo.FindByEmail(context.Background(), email)
	assert.NoError(t, err)
	assert.Equal(t, userOut.FirstName, firstName)
	assert.Equal(t, userOut.MiddleName, middleName)
	assert.Equal(t, userOut.LastName, lastName)
	assert.Equal(t, userOut.HashPassword, hashPass)
	assert.Equal(t, userOut.Email, email)
}

func TestFindByEmail_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	repo := NewPostgres(log, db)

	email := "mail@mail.ru"

	mock.ExpectQuery("SELECT id, first_name, middle_name, last_name, hash_password, email FROM users WHERE email = \\$1").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.FindByEmail(context.Background(), email)

	assert.ErrorIs(t, err, repository.ErrUserNotFound)
}
