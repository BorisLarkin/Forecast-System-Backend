package handler

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *Handler) WithAuthCheck(ctx *gin.Context) {
	// Извлечение токена из заголовка Authorization
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен"})
		ctx.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Парсинг и валидация токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return []byte(os.Getenv("JWT_KEY")), nil // Используем тот же секретный ключ
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
		ctx.Abort()
		return
	}

	// Извлечение userID из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка обработки токена"})
		ctx.Abort()
		return
	}

	userID := uint(claims["userID"].(float64))
	isModerator := claims["isModerator"].(bool)

	// Проверка сессии в Redis
	redisToken, err := h.Repository.GetSession(ctx.Request.Context(), userID)
	if err != nil || redisToken != tokenString {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Сессия не найдена или истекла"})
		ctx.Abort()
		return
	}

	// Передаем userID в контекст
	ctx.Set("userID", userID)
	ctx.Set("isModerator", isModerator)
	ctx.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *Handler) ModeratorMiddleware(ctx *gin.Context) {
	_, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authed"})
		ctx.Abort()
		return
	}

	isModerator := ctx.MustGet("isModerator").(bool)

	if !isModerator {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
