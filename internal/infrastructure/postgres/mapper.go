package postgres

import (
	userdomain "auth/internal/domain/user"
)

func modelToDomain(pm *PostgresUserModel) *userdomain.UserDomain {
	return userdomain.RestoreUserDomain(
		pm.FirstName,
		pm.MiddleName.String,
		pm.LastName,
		pm.HashPassword,
		pm.Email,
	)
}

func domainToModel(ud *userdomain.UserDomain) *PostgresUserModel {
	return NewPostgresModel(
		0,
		ud.FirstName,
		ud.MiddleName,
		ud.LastName,
		ud.HashPassword,
		ud.Email,
	)
}
