package jwt

import (
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/echo/config"
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

type JWTHandler struct {
	secretKey []byte
}

func NewJWTHandler(secretKey string) *JWTHandler {
	return &JWTHandler{
		secretKey: []byte(config.AppConfig.JWTConfig.SecretKey),
	}
}

func (j *JWTHandler) GenerateToken(userId uint) (string, error) {
	claims := customClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mosting",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString(j.secretKey)
}

func (j *JWTHandler) VerifyToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
