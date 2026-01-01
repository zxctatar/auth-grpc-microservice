package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustLoad_OK(t *testing.T) {
	cfgPath := "../../config/local.yaml"
	os.Setenv("CONFIG_PATH", cfgPath)
	defer os.Unsetenv("CONFIG_PATH")

	cfg := MustLoad()

	assert.Equal(t, cfg.GRPC.TimeOut, 15*time.Second)
	assert.Equal(t, cfg.GRPC.Port, 44044)
	assert.Equal(t, cfg.Postgres.Port, uint(5432))
	assert.Equal(t, cfg.Postgres.DbName, "users")
}
