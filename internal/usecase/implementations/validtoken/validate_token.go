package validtoken

import (
	"auth/internal/repository/tokenservice"
	"log/slog"
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

func (v *ValidateTokenUC) ValidateToken(token string) (uint32, error) {
	panic("not implemented")
}
