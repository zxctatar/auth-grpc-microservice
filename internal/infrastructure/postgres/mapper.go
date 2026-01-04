package postgres

import (
	userdomain "auth/internal/domain/user"
	posmodel "auth/internal/infrastructure/postgres/posmodels"
	logmodel "auth/internal/usecase/models/login"
)

func modelToDomain(pm *posmodel.PostgresUserModel) *userdomain.UserDomain {
	return userdomain.RestoreUserDomain(
		pm.FirstName,
		pm.MiddleName.String,
		pm.LastName,
		pm.HashPassword,
		pm.Email,
	)
}

func domainToModel(ud *userdomain.UserDomain) *posmodel.PostgresUserModel {
	return posmodel.NewPostgresModel(
		0,
		ud.FirstName,
		ud.MiddleName,
		ud.LastName,
		ud.HashPassword,
		ud.Email,
	)
}

func modelToUserAuthData(pu *posmodel.PostgresUserAuthDataModel) *logmodel.UserAuthData {
	return logmodel.NewUserAuthData(
		pu.Id,
		pu.HashPassword,
	)
}
