package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) CreateList(userId int, list todo.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := r.db.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListRelationQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = r.db.Exec(createUserListRelationQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, nil
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.List, error) {
	var lists []todo.List

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable,
	)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.List, error) {
	var list todo.List

	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl 
				INNER JOIN %s ul on tl.id = ul.list_id 
				WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable,
	)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s tl USING %s ul 
				WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable,
	)
	result, err := r.db.Exec(query, userId, listId)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("there is no todo list with such id")
	}

	return nil
}
