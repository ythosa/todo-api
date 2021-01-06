package dto

import (
	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

type ItemsResponse struct {
	Data []models.TodoItem `json:"data"`
}
