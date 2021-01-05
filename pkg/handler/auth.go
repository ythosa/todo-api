package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Inexpediency/todo-rest-api"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
