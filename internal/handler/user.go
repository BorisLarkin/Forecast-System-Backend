package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web/internal/ds"
	"web/internal/utils"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) LogoutUser(ctx *gin.Context) {
	if err := h.Repository.Logout(); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfuly logged out"})
}

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
