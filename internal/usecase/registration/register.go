package registration

import (
	"auth/internal/repository"
	"context"
	"log/slog"
)

type RegistrationUC struct {
	log *slog.Logger

	repo repository.StorageRepo
}

func NewRegistrationUC(log *slog.Logger, repo repository.StorageRepo) *RegistrationUC {
	return &RegistrationUC{
		log:  log,
		repo: repo,
	}
}

func (ru *RegistrationUC) RegUser(ctx context.Context, ri *RegInput) (uint32, error) {
	panic("not implemented")
}
