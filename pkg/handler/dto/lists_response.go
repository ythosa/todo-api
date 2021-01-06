package dto

import "github.com/Inexpediency/todo-rest-api"

type ListsResponse struct {
	Data []todo.List `json:"data"`
}
