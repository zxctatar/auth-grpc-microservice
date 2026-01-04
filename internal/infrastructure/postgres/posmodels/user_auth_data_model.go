package posmodel

type PostgresUserAuthDataModel struct {
	Id           uint32 `db:"id"`
	HashPassword string `db:"hash_password"`
}

func NewPostgresUserAuthDataModel(id uint32, hashPassword string) *PostgresUserAuthDataModel {
	return &PostgresUserAuthDataModel{
		Id:           id,
		HashPassword: hashPassword,
	}
}
