package repository

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item entity.TodoItem) (int, error)
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
