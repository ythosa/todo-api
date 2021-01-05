package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserByUsername(username string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
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
	}
}
