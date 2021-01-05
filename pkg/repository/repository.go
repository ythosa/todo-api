package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
