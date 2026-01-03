package registration

import (
	"auth/internal/repository"
	regmodels "auth/internal/usecase/models/registration"
	"context"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

var (
	invalidId uint32 = 0
)

type RegistrationUC struct {
	log  *slog.Logger
	repo repository.StorageRepo
}

func NewRegistrationUC(log *slog.Logger, repo repository.StorageRepo) *RegistrationUC {
	return &RegistrationUC{
		log:  log,
		repo: repo,
	}
}

func (ru *RegistrationUC) RegUser(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
	const op = "registration.RegUser"

	log := ru.log.With(slog.String("op", op), slog.String("email", ri.Email))

	log.Info("starting user registartion")

	hashPass, err := bcrypt.GenerateFromPassword([]byte(ri.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("failed to create hash password", slog.String("error", err.Error()))
		return invalidId, err
	}

	userInput, err := modelToDomain(ri, string(hashPass))

	if err != nil {
		log.Warn("failed to create a user domain from reg input", slog.String("error", err.Error()))
		return invalidId, err
	}

	userOut, err := ru.repo.FindByEmail(ctx, userInput.Email)

	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		log.Error("failed to find by email", slog.String("error", err.Error()))
		return invalidId, err
	}
	if userOut != nil {
		log.Info("user already exists")
		return invalidId, ErrUserAlreadyExists
	}

	newId, err := ru.repo.Save(ctx, userInput)

	if err != nil {
		log.Error("failed to save user", slog.String("error", err.Error()))
		return invalidId, err
	}

	return newId, err
}
