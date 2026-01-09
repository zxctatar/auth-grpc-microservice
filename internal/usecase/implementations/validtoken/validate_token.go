package validtoken

import (
	"auth/internal/repository/tokenservice"
	"context"
	"log/slog"
)

var (
	invalidId = uint32(0)
)

type ValidateTokenUC struct {
	log        *slog.Logger
	tokService tokenservice.TokenService
}

func NewValidateTokenUC(log *slog.Logger, tokService tokenservice.TokenService) *ValidateTokenUC {
	return &ValidateTokenUC{
		log:        log,
		tokService: tokService,
	}
}

func (v *ValidateTokenUC) ValidateToken(ctx context.Context, token string) (uint32, error) {
	const op = "validtoken.ValidateToken"

	log := v.log.With(slog.String("op", op))

	log.Info("starting token verification")

	id, err := v.tokService.ValidateToken(token)

	if err != nil {
		log.Warn("the token is invalid", slog.String("error", err.Error()))
		return invalidId, err
	}

	log.Info("token verification successful")

	return id, nil
}
