package handler

import (
	"fmt"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Role ds.Role
	Uid  uint
}

func (h *Handler) GetTokenPayload(gCtx *gin.Context) (*Payload, error) {
	jwtStr := gCtx.GetHeader("Authorization")
	if len(jwtStr) < len(jwtPrefix) {
		return nil, fmt.Errorf("no valid auth header provided")
	}
	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Config.JWT.Key), nil
	})
	if err != nil {
		return nil, err
	}

	myClaims := token.Claims.(*ds.JWTClaims)

	return &Payload{Role: myClaims.Role, Uid: myClaims.UserID}, nil
}
