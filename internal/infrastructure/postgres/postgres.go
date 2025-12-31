package postgres

import (
	userdomain "auth/internal/domain/user"
	"auth/internal/repository"
	"context"
	"database/sql"
	"errors"
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

func (p *Postgres) Save(ctx context.Context, user *userdomain.UserDomain) (uint32, error) {
	panic("not implemented")
}

func (p *Postgres) FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error) {
	const op = "postgres.FindByEmail"

	log := p.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("start find user by email")

	row := p.db.QueryRowContext(ctx, "SELECT id, first_name, middle_name, last_name, hash_password, email FROM users WHERE email = $1", email)

	var posModel PostgresModel

	err := row.Scan(
		&posModel.Id,
		&posModel.FirstName,
		&posModel.MiddleName,
		&posModel.LastName,
		&posModel.HashPassword,
		&posModel.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info("unsuccessful user search by email", slog.String("error", err.Error()))
			return nil, repository.ErrUserNotFound
		}
		log.Error("error searching for user by email", slog.String("error", err.Error()))
		return nil, err
	}

	userOut := modelToDomain(&posModel)

	p.log.Info("user found")

	return userOut, nil
}
