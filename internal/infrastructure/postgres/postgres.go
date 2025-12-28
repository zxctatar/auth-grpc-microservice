package postgres

import (
	userdomain "auth/internal/domain/user"
	"context"
	"database/sql"
	"log/slog"
)

type Postgres struct {
	log *slog.Logger
	db  *sql.DB
}

func NewPostgres(log *slog.Logger, db *sql.DB) *Postgres {
	return &Postgres{
		log: log,
		db:  db,
	}
}

func (p *Postgres) Save(ctx context.Context, user *userdomain.UserDomain) (int, error) {
	panic("not implemented")
}

func (p *Postgres) FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error) {
	panic("not implemented")
}
