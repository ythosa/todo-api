package service

import (
	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input dto.UpdateList) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
