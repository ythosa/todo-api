package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/handler/dto"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.List
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.GetAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
