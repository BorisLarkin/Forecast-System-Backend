package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"web/internal/ds"
	redis_api "web/internal/redis-api"
	"web/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

/*
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var req ds.Users

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}
	user, err := h.Repository.RegiterUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
*/

func (h *Handler) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var req ds.Users

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	err := h.Repository.UpdateUser(req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    userID,
		"login": req.Login,
		"role":  req.Role,
	})
}

/*
	func (h *Handler) LogoutUser(ctx *gin.Context) {
		if err := h.Repository.Logout(); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfuly logged out"})
	}
*/
type loginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
}

// Login godoc
// @Summary      Login the specified user
// @Description  very very friendly response
// @Tags         Userss
// @Produce      json
// @Success      200  {object}  loginResp
// @Router       /user/login [post]
func (h *Handler) LoginUser(gCtx *gin.Context) {
	req := &loginReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := h.Repository.GetUserByLogin(req.Login)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if req.Login == user.Login && req.Password == user.Password {
		token, err := utils.GenerateJWT(h.Config, user.User_id, ds.Role(user.Role))
		if err != nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to generate a token"))
			return
		}
		if token == nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))
			return
		}
		strToken, err := token.SignedString([]byte(h.Config.JWT.Token))

		if err != nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cant create str token"))
			return
		}

		gCtx.JSON(http.StatusOK, loginResp{
			ExpiresIn:   h.Config.JWT.ExpiresIn,
			AccessToken: strToken,
			TokenType:   "Bearer",
		})
	}

	gCtx.AbortWithStatus(http.StatusForbidden) // отдаем 403 ответ в знак того что доступ запрещен
}

type registerReq struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type registerResp struct {
	Ok bool `json:"ok"`
}

func (h *Handler) Register(gCtx *gin.Context) {
	req := &registerReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Pass == "" {
		gCtx.AbortWithError(http.StatusBadRequest, fmt.Errorf("pass is empty"))
		return
	}

	if req.Login == "" {
		gCtx.AbortWithError(http.StatusBadRequest, fmt.Errorf("login is empty"))
		return
	}

	err = h.Repository.RegisterUser(&ds.Users{
		Role:     int(ds.User),
		Login:    req.Login,
		Password: generateHashString(req.Pass), // пароли делаем в хешированном виде и далее будем сравнивать хеши, чтобы их не угнали с базой вместе
	})
	if err != nil {
		gCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gCtx.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (h *Handler) Logout(gCtx *gin.Context) {
	// получаем заголовок
	jwtStr := gCtx.GetHeader("Authorization")
	if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
		gCtx.AbortWithStatus(http.StatusBadRequest) // отдаем что нет доступа

		return // завершаем обработку
	}

	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	_, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Config.JWT.Token), nil
	})
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)

		return
	}

	// сохраняем в блеклист редиса
	err = redis_api.WriteJWTToBlacklist(h.Repository.RedisClient, gCtx.Request.Context(), jwtStr, h.Config.JWT.ExpiresIn)
	if err != nil {
		gCtx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	gCtx.Status(http.StatusOK)
}
