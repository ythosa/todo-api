package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/dto"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserByUsername(username string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.List) (int, error)
	GetAll(userId int) ([]todo.List, error)
	GetById(userId, listId int) (todo.List, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input dto.UpdateList) error
}

type TodoItem interface {
	Create(listId int, item todo.Item) (int, error)
	GetAll(userId, listId int) ([]todo.Item, error)
	GetById(userId, itemId int) (todo.Item, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input dto.UpdateItem) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository returns new repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
