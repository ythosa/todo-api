package service

import (
	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(listId int, item models.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input dto.UpdateItem) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}
