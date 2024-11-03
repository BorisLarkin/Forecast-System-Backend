package handler

import (
	"fmt"
	"net/http"
	"strings"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var input ds.Users

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if strings.TrimSpace(input.Login) == "" || strings.TrimSpace(input.Password) == "" {
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
		"id":           input.User_id,
		"login":        input.Login,
		"is_moderator": input.IsAdmin,
	})
}

func (h *Handler) AuthUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.Repository.Auth(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user, _ := h.Repository.GetUserByID(id)
	ctx.JSON(http.StatusAccepted, user)
}

func (h *Handler) DeAuthUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.Repository.Deauth(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfuly deauthed user %s", id)})
}
