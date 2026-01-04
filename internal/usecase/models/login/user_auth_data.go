package logmodel

type UserAuthData struct {
	Id           uint32
	HashPassword string
}

func NewUserAuthData(id uint32, hashPassword string) *UserAuthData {
	return &UserAuthData{
		Id:           id,
		HashPassword: hashPassword,
	}
}
