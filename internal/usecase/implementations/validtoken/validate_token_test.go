package validtoken

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidateTokenUC_Success(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	tokService := &tokenServiceMock{
		genFn: func(id uint32) (string, error) {
			t.Fatal("token should not be created when verifying the token")
			return "", nil
		},
		validTokenFn: func(token string) (uint32, error) {
			return 1, nil
		},
	}

	validTokenUc := NewValidateTokenUC(log, tokService)

	id, err := validTokenUc.ValidateToken(context.Background(), "token")

	assert.NoError(t, err)
	assert.Equal(t, uint32(1), id)
	assert.True(t, tokService.validTokenCalled)
}
