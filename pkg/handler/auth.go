package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input dto.SignIn

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid username or password")
		return
	}

	tokens, err := h.services.Authorization.GenerateTokens(input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) refreshTokens(c *gin.Context) {
	refreshToken := c.GetHeader("Refresh")
	if refreshToken == "" {
		newErrorResponse(c, http.StatusBadRequest, "Refresh token is not provided")
	}

	tokens, err := h.services.Authorization.RefreshTokens(refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}
