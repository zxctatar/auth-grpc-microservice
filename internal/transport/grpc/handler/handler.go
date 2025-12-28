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

func (ah *AuthHandler) Registration(context.Context, *authv1.RegistrationRequest) (*authv1.RegistrationResponse, error) {
	const op = "handler.Registration"

	ah.log.Info("new registration request", slog.String("op", op))

	panic("not implemented")
}

func (ah *AuthHandler) Login(context.Context, *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	const op = "handler.Login"

	ah.log.Info("new login request", slog.String("op", op))

	panic("not implemented")
}

func (ah *AuthHandler) ValidateToken(context.Context, *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	const op = "handler.ValidateToken"

	ah.log.Info("new validate token request", slog.String("op", op))

	panic("not implemented")
}
