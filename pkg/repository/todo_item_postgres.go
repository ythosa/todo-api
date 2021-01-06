package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item models.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		todoItemsTable,
	)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemRelationQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)")
	if _, err := r.db.Exec(createListItemRelationQuery, listId, itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, nil
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
				INNER JOIN %s li on li.item_id = ti.id
				INNER JOIN %s ul on ul.list_id = li.list_id 
				WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
				INNER JOIN %s li on li.item_id = ti.id
				INNER JOIN %s ul on ul.list_id = li.list_id 
				WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ti USING %s li, %s ul 
				WHERE ti.id = li.item_id AND li.id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)

	result, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("there is no todo item with such id")
	}

	return nil
}

func (r *TodoItemPostgres) Update(userId, itemId int, input dto.UpdateItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s ti SET %s FROM %s li, %s ul
				WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("there is no todo item with such id")
	}

	return nil
}
