package handler

import (
	userdomain "auth/internal/domain/user"
	authv1 "auth/internal/transport/grpc/pb"
	"auth/internal/usecase/implementations/login"
	"auth/internal/usecase/implementations/registration"
	usecaseinterf "auth/internal/usecase/interfaces"
	logmodel "auth/internal/usecase/models/login"
	regmodels "auth/internal/usecase/models/registration"
	"context"
	"errors"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer

	log     *slog.Logger
	timeOut *time.Duration
	regUC   usecaseinterf.RegistrationUseCase
	loginUC usecaseinterf.LoginUseCase
}

func NewAuthHandler(log *slog.Logger, timeOut *time.Duration, regUC usecaseinterf.RegistrationUseCase, loginUC usecaseinterf.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		log:     log,
		timeOut: timeOut,
		regUC:   regUC,
		loginUC: loginUC,
	}
}

func (ah *AuthHandler) Registration(ctx context.Context, rr *authv1.RegistrationRequest) (*authv1.RegistrationResponse, error) {
	const op = "handler.Registration"

	log := ah.log.With(slog.String("op", op), slog.String("email", rr.Email))

	log.Info("new registration request")

	ctx, cancel := context.WithTimeout(ctx, *ah.timeOut)
	defer cancel()

	regInput, err := regmodels.NewRegInput(
		rr.FirstName,
		rr.MiddleName,
		rr.LastName,
		rr.Password,
		rr.Email,
	)

	if err != nil {
		log.Warn("cannot to create reg input", slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := ah.regUC.RegUser(ctx, regInput)

	if err != nil {
		if errors.Is(err, registration.ErrUserAlreadyExists) {
			log.Info("registration failed", slog.String("error", err.Error()))
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, userdomain.ErrInvalidEmail) {
			log.Warn("registration failed", slog.String("error", err.Error()))
			return nil, status.Error(codes.InvalidArgument, "invalid mail format")
		} else if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Warn("long query execution", slog.String("error", err.Error()))
			return nil, status.Error(codes.Canceled, "the request was canceled")
		}
		log.Warn("unsuccessful user registration", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	log.Info("sending a response to the client")

	return &authv1.RegistrationResponse{
		UserId: id,
	}, nil
}

func (ah *AuthHandler) Login(ctx context.Context, lg *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	const op = "handler.Login"

	log := ah.log.With(slog.String("op", op), slog.String("email", lg.Email))

	log.Info("new login request")

	ctx, cancel := context.WithTimeout(ctx, *ah.timeOut)
	defer cancel()

	loginInput, err := logmodel.NewLoginInput(
		lg.Email,
		lg.Password,
	)

	if err != nil {
		ah.log.Warn("cannot to create login input", slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := ah.loginUC.Login(ctx, loginInput)

	if err != nil {
		if errors.Is(err, login.ErrUserNotFound) {
			log.Info("login failed", slog.String("error", err.Error()))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, login.ErrWrongPassword) {
			log.Info("login failed", slog.String("error", err.Error()))
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		log.Warn("unsuccessful user login", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &authv1.LoginResponse{
		Token: token,
	}, nil
}

func (ah *AuthHandler) ValidateToken(ctx context.Context, vtr *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	const op = "handler.ValidateToken"

	ah.log.Info("new validate token request", slog.String("op", op))

	panic("not implemented")
}
