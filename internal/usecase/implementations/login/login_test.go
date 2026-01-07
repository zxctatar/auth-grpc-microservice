package login

import (
	"auth/internal/repository/storagerepo"
	logmodel "auth/internal/usecase/models/login"
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	repoMock := &storageRepoMock{
		findAuthDataByEmail: func(ctx context.Context, email string) (*logmodel.UserAuthData, error) {
			assert.Equal(t, "mail@mail.ru", email)
			return logmodel.NewUserAuthData(1, "somePass"), nil
		},
	}
	tokenMock := &tokenServiceMock{
		genFn: func(id uint32) (string, error) {
			assert.Equal(t, uint32(1), id)
			return "token", nil
		},
	}
	hashMock := &hashServiceMock{
		comFn: func(hashPass, password []byte) error {
			assert.Equal(t, []byte("somePass"), password)
			return nil
		},
	}

	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	login := NewLoginUc(log, repoMock, tokenMock, hashMock)

	email := "mail@mail.ru"
	passWord := "somePass"

	logInput, err := logmodel.NewLoginInput(email, passWord)

	assert.NoError(t, err)

	token, err := login.Login(context.Background(), logInput)

	assert.NoError(t, err)
	assert.Equal(t, "token", token)

	assert.True(t, repoMock.findAuthDataByEmailCalled)
	assert.True(t, tokenMock.genFnCalled)
	assert.True(t, hashMock.comCalled)
}

func TestLogin_UserNotFound(t *testing.T) {
	repoMock := &storageRepoMock{
		findAuthDataByEmail: func(ctx context.Context, email string) (*logmodel.UserAuthData, error) {
			return nil, storagerepo.ErrUserNotFound
		},
	}
	tokenMock := &tokenServiceMock{
		genFn: func(id uint32) (string, error) {
			t.Fatal("token must not be generated if user not found")
			return "", nil
		},
	}
	hashMock := &hashServiceMock{
		comFn: func(hashPass, password []byte) error {
			t.Fatal("password must not be compared if user not found")
			return nil
		},
	}

	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	login := NewLoginUc(log, repoMock, tokenMock, hashMock)

	email := "mail@mail.ru"
	passWord := "somePass"

	logInput, err := logmodel.NewLoginInput(email, passWord)

	assert.NoError(t, err)

	_, err = login.Login(context.Background(), logInput)

	assert.ErrorIs(t, err, storagerepo.ErrUserNotFound)

	assert.True(t, repoMock.findAuthDataByEmailCalled)
}
