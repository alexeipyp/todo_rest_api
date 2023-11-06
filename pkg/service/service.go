package service

import (
	todo "github.com/alexeipyp/todo_rest_api"
	"github.com/alexeipyp/todo_rest_api/pkg/repository"
	"github.com/alexeipyp/todo_rest_api/pkg/service/auth"
	todolist "github.com/alexeipyp/todo_rest_api/pkg/service/todo_list"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func New(repos *repository.Repository) *Service {
	return &Service{
		Authorization: auth.New(repos.Authorization),
		TodoList:      todolist.New(repos.TodoList),
	}
}
