package service

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/Konstantsiy/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
