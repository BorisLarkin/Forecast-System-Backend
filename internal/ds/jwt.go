package ds

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims      // все что точно необходимо по RFC
	UserID             uint `json:"user_uuid"` // наши данные - uuid этого пользователя в базе данных
	Role               Role `json:"role"`      // список доступов в нашей системе
}
