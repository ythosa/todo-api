package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	return 0, nil
}
