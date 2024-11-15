package utils

import (
	"time"
	"web/internal/config"
	"web/internal/ds"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(cfg *config.Config, userID uint, role ds.Role) (string, error) {
	token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "weather-admin",
		},
		UserID: userID,
		Role:   role,
	})
	tokenstr, err := token.SignedString([]byte(cfg.JWT.Key))
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}
