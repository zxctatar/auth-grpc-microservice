package login

import (
	"auth/internal/repository/storagerepo"
	"auth/internal/repository/tokenservice"
	logmodel "auth/internal/usecase/models/login"
	"context"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

var (
	invalidToken = ""
)

type LoginUC struct {
	log        *slog.Logger
	repo       storagerepo.StorageRepo
	tokService tokenservice.TokenService
}

func NewLoginUc(log *slog.Logger, repo storagerepo.StorageRepo, tokService tokenservice.TokenService) *LoginUC {
	return &LoginUC{
		log:        log,
		repo:       repo,
		tokService: tokService,
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

	err = bcrypt.CompareHashAndPassword([]byte(authData.HashPassword), []byte(li.Password))

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
