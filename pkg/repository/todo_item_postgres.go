package repository

import (
	"fmt"
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item entity.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id")
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createItemsListQuery := fmt.Sprintf("insert into %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createItemsListQuery, listId, itemId)
	return itemId, tx.Commit()
}
