package handler

import (
	"net/http"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var input ds.Users

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if input.Login == "" || input.Password == "" {
		ctx.JSON(http.StatusBadRequest, "Login and password are required")
		return
	}

	if err := h.Repository.CreateUser(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       input.User_id,
		"login":    input.Login,
		"is_admin": input.IsAdmin,
	})
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var input ds.Users

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	err := h.Repository.UpdateUser(input, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":           userID,
		"login":        input.Login,
		"is_moderator": input.IsAdmin,
	})
}

func (h *Handler) AuthUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Repository.Auth(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.Repository.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user.User_id, "login": user.Login, "is_admin": user.IsAdmin})
}

func (h *Handler) DeAuthUser(ctx *gin.Context) {
	if err := h.Repository.Deauth(); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfuly deauthed user"})
}
