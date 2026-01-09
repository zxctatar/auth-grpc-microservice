package jwtservice

import (
	"auth/internal/repository/tokenservice"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	invalidToken = ""
	invalidId    = uint32(0)
)

type CustomClaim struct {
	Id uint32 `json:"id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	SecretKey []byte
	TimeOut   *time.Duration
}

func NewJWTService(secretKey []byte, timeOut *time.Duration) *JWTService {
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
	tokenString, err := token.SignedString(j.SecretKey)

	if err != nil {
		return invalidToken, err
	}

	return tokenString, nil
}

func (j *JWTService) ValidateToken(token string) (uint32, error) {
	claim := &CustomClaim{}

	tok, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.SecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return invalidId, tokenservice.ErrInvalidSignature
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return invalidId, tokenservice.ErrTokenMalformed
		}
		return invalidId, err
	}

	if !tok.Valid {
		return invalidId, ErrInvalidToken
	}

	return claim.Id, nil
}
