package posmodel

import "database/sql"

type PostgresUserModel struct {
	Id           uint32         `db:"id"`
	FirstName    string         `db:"first_name"`
	MiddleName   sql.NullString `db:"middle_name"`
	LastName     string         `db:"last_name"`
	HashPassword string         `db:"hash_password"`
	Email        string         `db:"email"`
}

func NewPostgresModel(id uint32, firstName, middleName, lastName, hashPassword, email string) *PostgresUserModel {
	midName := sql.NullString{
		String: middleName,
		Valid:  middleName != "",
	}

	return &PostgresUserModel{
		Id:           id,
		FirstName:    firstName,
		MiddleName:   midName,
		LastName:     lastName,
		HashPassword: hashPassword,
		Email:        email,
	}
}
