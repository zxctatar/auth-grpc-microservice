package postgres

import (
	userdomain "auth/internal/domain/user"
	posmodel "auth/internal/infrastructure/postgres/posmodels"
	"auth/internal/repository/storagerepo"
	logmodel "auth/internal/usecase/models/login"
	"context"
	"database/sql"
	"errors"
	"log/slog"
)

var (
	invalidId uint32 = 0
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
	const op = "postgres.Save"

	log := p.log.With(slog.String("op", op), slog.String("user email", user.Email))

	log.Info("start save user")

	posModel := domainToModel(user)

	row := p.db.QueryRowContext(ctx,
		`INSERT INTO users (first_name, middle_name, last_name, hash_password, email)
	VALUES($1, $2, $3, $4, $5) RETURNING id`,
		posModel.FirstName,
		posModel.MiddleName,
		posModel.LastName,
		posModel.HashPassword,
		posModel.Email,
	)

	var userId uint32 = 0

	err := row.Scan(&userId)

	if err != nil {
		log.Warn("failed to save user", slog.String("error", err.Error()))
		return invalidId, err
	}

	log.Info("the user has been saved")

	return userId, nil
}

func (p *Postgres) FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error) {
	const op = "postgres.FindByEmail"

	log := p.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("start find user by email")

	row := p.db.QueryRowContext(ctx, "SELECT id, first_name, middle_name, last_name, hash_password, email FROM users WHERE email = $1", email)

	var posModel posmodel.PostgresUserModel

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
			return nil, storagerepo.ErrUserNotFound
		}
		log.Error("error searching for user by email", slog.String("error", err.Error()))
		return nil, err
	}

	userOut := modelToDomain(&posModel)

	log.Info("user found")

	return userOut, nil
}

func (p *Postgres) FindAuthDataByEmail(ctx context.Context, email string) (*logmodel.UserAuthData, error) {
	const op = "postgres.FindAuthDataByEmail"

	log := p.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("start find auth data by email")

	row := p.db.QueryRowContext(ctx, "SELECT id, hash_password FROM users WHERE email = $1", email)

	var posAuthData posmodel.PostgresUserAuthDataModel

	err := row.Scan(
		&posAuthData.Id,
		&posAuthData.HashPassword,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info("unsuccessful attempt to get auth user data", slog.String("error", err.Error()))
			return nil, storagerepo.ErrUserNotFound
		}
		log.Error("error retrieving auth user data", slog.String("error", err.Error()))
		return nil, err
	}

	userAuthData := modelToUserAuthData(&posAuthData)

	log.Info("data found")

	return userAuthData, nil
}
