package utils

import (
	"time"
	"web/internal/config"
	"web/internal/ds"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(cfg *config.Config, userID uint, role ds.Role) (*jwt.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ds.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bitop-admin",
		},
		UserID: userID,
		Role:   role,
	})
	return token, nil
}
