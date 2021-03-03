package service

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/Konstantsiy/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item entity.TodoItem) (int, error)
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
