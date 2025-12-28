package handler

import (
	authv1 "auth/internal/transport/grpc/pb"
	"auth/internal/usecase/registration"
	"context"
	"log/slog"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer

	log   *slog.Logger
	regUC *registration.RegistrationUC
}

func NewAuthHandler(log *slog.Logger, regUC *registration.RegistrationUC) *AuthHandler {
	return &AuthHandler{
		log:   log,
		regUC: regUC,
	}
}

func (ah *AuthHandler) Registration(ctx context.Context, rr *authv1.RegistrationRequest) (*authv1.RegistrationResponse, error) {
	const op = "handler.Registration"

	ah.log.Info("new registration request", slog.String("op", op))

	regInput, err := registration.NewRegInput(
		rr.FirstName,
		rr.MiddleName,
		rr.LastName,
		rr.Password,
		rr.Email,
	)

	if err != nil {
		ah.log.Warn("cannot to create reg input", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	id, err := ah.regUC.RegUser(ctx, regInput)

	if err != nil {
		ah.log.Warn("unsuccessful user registration", slog.String("op", op), slog.String("error", err.Error()))
		return nil, err
	}

	return &authv1.RegistrationResponse{
		UserId: uint32(id),
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
