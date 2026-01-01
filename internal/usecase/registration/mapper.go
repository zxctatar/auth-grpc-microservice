package registration

import (
	userdomain "auth/internal/domain/user"
	"auth/internal/usecase"
)

func modelToDomain(ri *usecase.RegInput, hashPass string) (*userdomain.UserDomain, error) {
	user, err := userdomain.NewUserDomain(
		ri.FirstName,
		ri.MiddleName,
		ri.LastName,
		hashPass,
		ri.Email,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
