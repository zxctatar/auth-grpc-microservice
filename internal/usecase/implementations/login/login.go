package login

import (
	"auth/internal/repository/hashservice"
	"auth/internal/repository/storagerepo"
	"auth/internal/repository/tokenservice"
	logmodel "auth/internal/usecase/models/login"
	"context"
	"errors"
	"log/slog"
)

var (
	invalidToken = ""
)

type LoginUC struct {
	log        *slog.Logger
	repo       storagerepo.StorageRepo
	tokService tokenservice.TokenService
	hasher     hashservice.HashService
}

func NewLoginUc(log *slog.Logger, repo storagerepo.StorageRepo, tokService tokenservice.TokenService, hasher hashservice.HashService) *LoginUC {
	return &LoginUC{
		log:        log,
		repo:       repo,
		tokService: tokService,
		hasher:     hasher,
	}
}

func (l *LoginUC) Login(ctx context.Context, li *logmodel.LoginInput) (string, error) {
	const op = "login.Login"

	log := l.log.With(slog.String("op", op), slog.String("email", li.Email))

	log.Info("starting user login")

	authData, err := l.repo.FindAuthDataByEmail(ctx, li.Email)

	if err != nil {
		if errors.Is(err, storagerepo.ErrUserNotFound) {
			log.Info("failed user login", slog.String("error", err.Error()))
			return invalidToken, err
		}
		log.Warn("failed user login", slog.String("error", err.Error()))
		return invalidToken, err
	}

	err = l.hasher.ComparePassword([]byte(authData.HashPassword), []byte(li.Password))

	if err != nil {
		log.Info("wrong password", slog.String("error", err.Error()))
		return invalidToken, ErrWrongPassword
	}

	token, err := l.tokService.Generate(authData.Id)

	if err != nil {
		log.Warn("unsuccessful token generation", slog.String("error", err.Error()))
		return invalidToken, err
	}

	return token, nil
}
