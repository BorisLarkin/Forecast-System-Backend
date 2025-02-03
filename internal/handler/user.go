package handler

import (
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

type updateReq struct {
	Login    string  `json:"login" binding:"required"`
	Role     ds.Role `json:"role" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

type updateResp struct {
	Uid   string  `json:"uid" binding:"required"`
	Login string  `json:"login" binding:"required"`
	Role  ds.Role `json:"role" binding:"required"`
}

// Update godoc
// @Summary      Update the specified user
// @Description  very very friendly response
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        user body updateReq true "New user data"
// @Param        Authorization header string true "Auth Bearer token header"
// @Success      200  {object}  updateResp
// @Failure      500
// @Router       /user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var updateReq updateReq

	if err := ctx.BindJSON(&updateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	err = h.Repository.UpdateUser(ds.Users{Login: updateReq.Login, Password: updateReq.Password, Role: int(updateReq.Role)}, userID, payload.Uid, payload.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updateResp{
		Uid:   userID,
		Login: updateReq.Login,
		Role:  ds.Role(updateReq.Role),
	})
}

type loginReq struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Guest    bool   `json:"guest" binding:"required"`
}

type loginResp struct {
	Login       string `json:"login" binding:"required"`
	Role        int    `json:"role" binding:"required"`
	ExpiresIn   string `json:"expires_in" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
	TokenType   string `json:"token_type" binding:"required"`
}

// Login godoc
// @Summary      Login the specified user
// @Description  very very friendly response
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body loginReq true "user data"
// @Success      200  {object} loginResp
// @Failure      403
// @Router       /user/login [post]
func (h *Handler) LoginUser(gCtx *gin.Context) {
	req := &loginReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if req.Guest {
		token, err := utils.GenerateJWT(h.Config, 0, ds.Role(ds.Guest))
		if err != nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to generate a token"))
			return
		}
		if token == "" {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))
			return
		}

		gCtx.JSON(http.StatusOK, loginResp{
			Login:       "Гость",
			Role:        0,
			ExpiresIn:   time.Duration(h.Config.JWT.ExpiresIn).String(),
			AccessToken: token,
			TokenType:   "Bearer",
		})
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
		if token == "" {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))
			return
		}

		gCtx.JSON(http.StatusOK, loginResp{
			Login:       user.Login,
			Role:        user.Role,
			ExpiresIn:   time.Duration(h.Config.JWT.ExpiresIn).String(),
			AccessToken: token,
			TokenType:   "Bearer",
		})
		return
	}

	gCtx.AbortWithStatus(http.StatusForbidden) // отдаем 403 ответ в знак того что доступ запрещен
}

type registerReq struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerResp struct {
	Ok bool `json:"ok" binding:"required"`
}

// Register godoc
// @Summary      Register the specified user
// @Description  very very friendly response
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body registerReq true "New user data"
// @Success      200  {object}  registerResp
// @Failure      400
// @Router       /user/register [post]
func (h *Handler) Register(gCtx *gin.Context) {
	req := &registerReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Password == "" {
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
		Password: req.Password, // пароли делаем в хешированном виде и далее будем сравнивать хеши, чтобы их не угнали с базой вместе
	})
	if err != nil {
		gCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gCtx.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

// Logout godoc
// @Summary      Logout the current user
// @Description  very very friendly response
// @Tags         Users
// @Produce      json
// @Param        Authorization header string true "Auth Bearer token header"
// @Success      200
// @Failure      500
// @Router       /user/logout [post]
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
		return []byte(h.Config.JWT.Key), nil
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
