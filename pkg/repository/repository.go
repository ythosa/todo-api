package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByUsername(username string) (models.User, error)
	SaveRefreshToken(userId int, token string) error
	GetRefreshToken(userId int) (string, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input dto.UpdateList) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input dto.UpdateItem) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository returns new repository
func NewRepository(db *sqlx.DB, cache *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db, cache),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
