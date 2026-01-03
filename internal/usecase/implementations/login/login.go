package login

import (
	"auth/internal/repository"
	logmodel "auth/internal/usecase/models/login"
	"context"
	"log/slog"
)

type LoginUC struct {
	log  *slog.Logger
	repo repository.StorageRepo
}

func NewLoginUc(log *slog.Logger, repo repository.StorageRepo) *LoginUC {
	return &LoginUC{
		log:  log,
		repo: repo,
	}
}

func (l *LoginUC) Login(ctx context.Context, li *logmodel.LoginInput) (string, error) {
	panic("not implemented")
}
