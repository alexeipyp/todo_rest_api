package service

import "github.com/alexeipyp/todo_rest_api/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func New(repos *repository.Repository) *Service {
	return &Service{}
}
