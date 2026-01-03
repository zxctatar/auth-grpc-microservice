package handler

import (
	userdomain "auth/internal/domain/user"
	authv1 "auth/internal/transport/grpc/pb"
	"auth/internal/usecase/implementations/registration"
	regmodels "auth/internal/usecase/models/registration"
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewAuthHandler_Succes(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	var regUCMock *registrationUCMock = nil

	handler := NewAuthHandler(log, &timeOut, regUCMock, nil)

	assert.Equal(t, handler.log, log)
	assert.Equal(t, handler.timeOut.Seconds(), timeOut.Seconds())
	assert.Equal(t, handler.regUC, regUCMock)
}

func TestRegistration_Success(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			return 1, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, nil)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	req := &authv1.RegistrationRequest{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   pass,
		Email:      email,
	}

	res, err := handler.Registration(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, res.UserId, uint32(1))
}

func TestRegistration_UserAlreadyExists(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			return 0, registration.ErrUserAlreadyExists
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, nil)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	req := &authv1.RegistrationRequest{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   pass,
		Email:      email,
	}

	res, err := handler.Registration(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.AlreadyExists)
}

func TestRegistration_InvalidEmail(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			return 0, userdomain.ErrInvalidEmail
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, nil)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "123"

	req := &authv1.RegistrationRequest{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   pass,
		Email:      email,
	}

	res, err := handler.Registration(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}

func TestRegistration_Canceled(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Millisecond
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			select {
			case <-time.After(2 * time.Millisecond):
				return 1, nil
			case <-ctx.Done():
				return 1, ctx.Err()
			}
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, nil)

	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "123"

	req := &authv1.RegistrationRequest{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   pass,
		Email:      email,
	}

	res, err := handler.Registration(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	if st.Code() != codes.DeadlineExceeded {
		assert.Equal(t, st.Code(), codes.Canceled)
	} else {
		assert.Equal(t, st.Code(), codes.DeadlineExceeded)
	}
}
