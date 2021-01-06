package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/handler/dto"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.Item) (int, error) {
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

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.Item, error) {

}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.Item, error) {

}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {

}

func (r *TodoItemPostgres) Update(userId, itemId int, input dto.UpdateItem) error {

}
