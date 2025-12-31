package postgres

import (
	userdomain "auth/internal/domain/user"
)

func modelToDomain(pm *PostgresModel) *userdomain.UserDomain {
	return userdomain.RestoreUserDomain(
		pm.FirstName,
		pm.MiddleName.String,
		pm.LastName,
		pm.HashPassword,
		pm.Email,
	)
}

func domainToModel(ud *userdomain.UserDomain) *PostgresModel {
	return NewPostgresModel(
		0,
		ud.FirstName,
		ud.MiddleName,
		ud.LastName,
		ud.HashPassword,
		ud.Email,
	)
}
