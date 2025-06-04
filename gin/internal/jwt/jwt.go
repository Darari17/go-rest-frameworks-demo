package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/helpers"
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

type JWTConfig struct {
	Jwt struct {
		SecretKey string `yaml:"secret_key"`
	} `yaml:"jwt"`
}

type JWTHandler struct {
	secretKey []byte
}

func NewJWTHandler(path string) *JWTHandler {
	var jwtCfg JWTConfig

	err := helpers.LoadYAMLConfig(path, &jwtCfg)
	if err != nil {
		log.Fatalf("failed to load jwt config: %v", err)
	}

	return &JWTHandler{
		secretKey: []byte(jwtCfg.Jwt.SecretKey),
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

	return claims, nil
}
