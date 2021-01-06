package service

import (
	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.List) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.List, error) {
	return s.repo.GetAll(userId)
}
