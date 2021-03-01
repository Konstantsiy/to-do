package repository

import (
	"fmt"
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list entity.TodoList) (int, error) {
	tx, err := r.db.Begin() // create transaction
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUserListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]entity.TodoList, error) {
	var lists []entity.TodoList
	getAllListsQuery := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id=$1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, getAllListsQuery, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (entity.TodoList, error) {
	var list entity.TodoList
	getListByIdQuery := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id=$1 and ul.list_id=$2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, getListByIdQuery, userId, listId)

	return list, err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	deleteQuery := fmt.Sprintf("delete form %s tl using %s ul where tl.id = ul.list_id and ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)
	_, err := r.db.Exec(deleteQuery, userId, listId)
	return err
}
