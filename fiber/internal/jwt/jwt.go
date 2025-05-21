package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	ID uint
	jwt.RegisteredClaims
}

type jwtHandler struct {
	secretKey []byte
}

type JWTHandler interface {
	CreateTokenJWT(id uint) (string, error)
	VerifyTokenJWT(tokenString string) (*jwtClaims, error)
}

func NewJWTHandler(secretKey string) JWTHandler {
	return &jwtHandler{
		secretKey: []byte(secretKey),
	}
}

// CreateTokenJWT implements JWTHandler.
func (j *jwtHandler) CreateTokenJWT(id uint) (string, error) {
	claims := jwtClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mosting",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString(j.secretKey)
}

// VerifiyTokenJWT implements JWTHandler.
func (j *jwtHandler) VerifyTokenJWT(tokenString string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
