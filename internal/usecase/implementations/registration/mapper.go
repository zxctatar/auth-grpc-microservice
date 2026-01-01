package registration

import (
	userdomain "auth/internal/domain/user"
	regmodels "auth/internal/usecase/models/registration"
)

func modelToDomain(ri *regmodels.RegInput, hashPass string) (*userdomain.UserDomain, error) {
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
