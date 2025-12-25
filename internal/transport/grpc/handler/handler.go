package handler

import (
	authv1 "auth/internal/transport/grpc/pb"
	"log/slog"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer

	log *slog.Logger
}

func NewAuthHandler(log *slog.Logger) *AuthHandler {
	return &AuthHandler{
		log: log,
	}
}
