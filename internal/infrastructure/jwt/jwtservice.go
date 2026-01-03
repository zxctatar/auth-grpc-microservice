package jwtservice

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	invalidToken = ""
)

type CustomClaim struct {
	Id uint32 `json:"id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	SecretKey string
	TimeOut   *time.Duration
}

func NewJWTService(secretKey string, timeOut *time.Duration) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
		TimeOut:   timeOut,
	}
}

func (j *JWTService) Generate(id uint32) (string, error) {
	expTime := time.Now().Add(*j.TimeOut)

	claim := &CustomClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "userId",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return invalidToken, err
	}

	return tokenString, nil
}
