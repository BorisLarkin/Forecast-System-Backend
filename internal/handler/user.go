package handler

import (
	"net/http"
	"strconv"
	"web/internal/ds"

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

func (h *Handler) LoginUser(ctx *gin.Context) {
	var req struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Login == "" || req.Password == "" {
		ctx.JSON(http.StatusBadRequest, "Login and password are required")
		return
	}
	id, err := h.Repository.Login(req.Login, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userid := strconv.Itoa(id)
	user, err := h.Repository.GetUserByID(userid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user.User_id, "login": user.Login, "is_admin": user.Role})
}

func (h *Handler) LogoutUser(ctx *gin.Context) {
	if err := h.Repository.Logout(); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfuly logged out"})
}
