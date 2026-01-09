package handler

import (
	userdomain "auth/internal/domain/user"
	"auth/internal/repository/storagerepo"
	"auth/internal/repository/tokenservice"
	authv1 "auth/internal/transport/grpc/pb"
	loginerror "auth/internal/usecase/errors/login"
	regerror "auth/internal/usecase/errors/registration"
	logmodel "auth/internal/usecase/models/login"
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
	var logUCMock *loginUCMock = nil
	var validateUCMock *validateTokenUCMock = nil

	handler := NewAuthHandler(log, &timeOut, regUCMock, logUCMock, validateUCMock)

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
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login usecase should not be called during registration")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during registration")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

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
	assert.True(t, regUCMock.regUserFnCalled)
}

func TestRegistration_UserAlreadyExists(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			return 0, regerror.ErrUserAlreadyExists
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login usecase should not be called during registration")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during registration")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

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
	assert.True(t, regUCMock.regUserFnCalled)
}

func TestRegistration_InvalidEmail(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 5 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			return 0, userdomain.ErrInvalidEmail
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login usecase should not be called during registration")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during registration")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

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
	assert.True(t, regUCMock.regUserFnCalled)
}

func TestRegistration_Canceled(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Millisecond
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			<-ctx.Done()
			return 1, ctx.Err()
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login usecase should not be called during registration")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during registration")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

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

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	res, err := handler.Registration(ctx, req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	if st.Code() != codes.DeadlineExceeded {
		assert.Equal(t, st.Code(), codes.Canceled)
	} else {
		assert.Equal(t, st.Code(), codes.DeadlineExceeded)
	}
	assert.True(t, regUCMock.regUserFnCalled)
}

func TestLogin_Success(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during login")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			return "token", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during login")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	email := "mail@mail.ru"
	password := "somePass"

	req := &authv1.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := handler.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, res.Token, "token")
	assert.True(t, loginUCMock.logFnCalled)
}

func TestLogin_UserNotFound(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during login")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			return "", storagerepo.ErrUserNotFound
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during login")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	email := "mail@mail.ru"
	password := "somePass"

	req := &authv1.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := handler.Login(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.NotFound)
	assert.True(t, loginUCMock.logFnCalled)
}

func TestLogin_WrongPassword(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during login")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			return "", loginerror.ErrWrongPassword
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during login")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	email := "mail@mail.ru"
	password := "somePass"

	req := &authv1.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := handler.Login(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.Unauthenticated)
	assert.True(t, loginUCMock.logFnCalled)
}

func TestLogin_Canceled(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Millisecond
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during login")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			<-ctx.Done()
			return "", ctx.Err()
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			t.Fatal("token validation use case should not be called during login")
			return 0, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	email := "mail@mail.ru"
	password := "somePass"

	req := &authv1.LoginRequest{
		Email:    email,
		Password: password,
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	res, err := handler.Login(ctx, req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	if st.Code() != codes.DeadlineExceeded {
		assert.Equal(t, st.Code(), codes.Canceled)
	} else {
		assert.Equal(t, st.Code(), codes.DeadlineExceeded)
	}
	assert.True(t, loginUCMock.logFnCalled)
}

func TestValidateToken_Success(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during token validation")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login use case should not be called during token validation")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			return 1, nil
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	token := "someToken"

	req := &authv1.ValidateTokenRequest{
		Token: token,
	}

	res, err := handler.ValidateToken(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, uint32(1), res.UserId)
	assert.True(t, res.Valid)
	assert.True(t, validateUCMock.validateFnCalled)
}

func TestValidateToken_InvalidSignature(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during token validation")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login use case should not be called during token validation")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			return 0, tokenservice.ErrInvalidSignature
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	token := "someToken"

	req := &authv1.ValidateTokenRequest{
		Token: token,
	}

	res, err := handler.ValidateToken(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.InvalidArgument)
	assert.True(t, validateUCMock.validateFnCalled)
}

func TestValidateToken_ErrTokenMalformed(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Second
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during token validation")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login use case should not be called during token validation")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			return 0, tokenservice.ErrTokenMalformed
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	token := "someToken"

	req := &authv1.ValidateTokenRequest{
		Token: token,
	}

	res, err := handler.ValidateToken(context.Background(), req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, st.Code(), codes.InvalidArgument)
	assert.True(t, validateUCMock.validateFnCalled)
}

func TestValidateToken_Canceled(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeOut := 1 * time.Millisecond
	regUCMock := &registrationUCMock{
		regUserFn: func(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
			t.Fatal("registration use case should not be called during token validation")
			return 0, nil
		},
	}
	loginUCMock := &loginUCMock{
		logFn: func(ctx context.Context, li *logmodel.LoginInput) (string, error) {
			t.Fatal("login use case should not be called during token validation")
			return "", nil
		},
	}
	validateUCMock := &validateTokenUCMock{
		validateFn: func(ctx context.Context, token string) (uint32, error) {
			<-ctx.Done()
			return 0, ctx.Err()
		},
	}

	handler := NewAuthHandler(log, &timeOut, regUCMock, loginUCMock, validateUCMock)

	token := "someToken"

	req := &authv1.ValidateTokenRequest{
		Token: token,
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	res, err := handler.ValidateToken(ctx, req)

	assert.Nil(t, res)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	if st.Code() != codes.DeadlineExceeded {
		assert.Equal(t, st.Code(), codes.Canceled)
	} else {
		assert.Equal(t, st.Code(), codes.DeadlineExceeded)
	}
	assert.True(t, validateUCMock.validateFnCalled)
}
