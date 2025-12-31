package handler

import (
	authv1 "auth/internal/transport/grpc/pb"
	"auth/internal/usecase/registration"
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
	regUC   *registration.RegistrationUC
}

func NewAuthHandler(log *slog.Logger, timeOut *time.Duration, regUC *registration.RegistrationUC) *AuthHandler {
	return &AuthHandler{
		log:     log,
		timeOut: timeOut,
		regUC:   regUC,
	}
}

func (ah *AuthHandler) Registration(ctx context.Context, rr *authv1.RegistrationRequest) (*authv1.RegistrationResponse, error) {
	const op = "handler.Registration"

	log := ah.log.With(slog.String("op", op))

	log.Info("new registration request")

	ctx, cancel := context.WithTimeout(ctx, *ah.timeOut)
	defer cancel()

	regInput, err := registration.NewRegInput(
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
			log.Info("registration is not possible", slog.String("error", err.Error()))
			return nil, status.Error(codes.AlreadyExists, err.Error())
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

	ah.log.Info("new login request", slog.String("op", op))

	panic("not implemented")
}

func (ah *AuthHandler) ValidateToken(ctx context.Context, vtr *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	const op = "handler.ValidateToken"

	ah.log.Info("new validate token request", slog.String("op", op))

	panic("not implemented")
}
