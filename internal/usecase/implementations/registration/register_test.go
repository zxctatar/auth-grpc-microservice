package registration

import (
	userdomain "auth/internal/domain/user"
	"auth/internal/repository/storagerepo"
	regmodels "auth/internal/usecase/models/registration"
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	repoMock := &storageRepoMock{
		findByEmailFn: func(ctx context.Context, email string) (*userdomain.UserDomain, error) {
			return nil, storagerepo.ErrUserNotFound
		},
		saveFn: func(ctx context.Context, user *userdomain.UserDomain) (uint32, error) {
			return 1, nil
		},
	}

	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	regUc := NewRegistrationUC(log, repoMock)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	ri, err := regmodels.NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.NoError(t, err)

	id, err := regUc.RegUser(context.Background(), ri)

	assert.NoError(t, err)
	assert.Equal(t, id, uint32(1))
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	repoMock := &storageRepoMock{
		findByEmailFn: func(ctx context.Context, email string) (*userdomain.UserDomain, error) {
			firstName := "Ivan"
			middleName := "Ivanovich"
			lastName := "Ivanov"
			hashPass := "somePass"

			user := userdomain.RestoreUserDomain(
				firstName,
				middleName,
				lastName,
				hashPass,
				email,
			)

			return user, nil
		},
		saveFn: func(ctx context.Context, user *userdomain.UserDomain) (uint32, error) {
			return invalidId, nil
		},
	}

	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	regUc := NewRegistrationUC(log, repoMock)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	ri, err := regmodels.NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.NoError(t, err)

	_, err = regUc.RegUser(context.Background(), ri)

	assert.ErrorIs(t, err, ErrUserAlreadyExists)
}
