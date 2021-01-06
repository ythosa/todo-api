package dto

import "github.com/Inexpediency/todo-rest-api"

type GetAllListsResponse struct {
	Data []todo.List `json:"data"`
}
