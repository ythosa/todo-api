package dto

import (
	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

type ListsResponse struct {
	Data []models.TodoList `json:"data"`
}
