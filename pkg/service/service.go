package service

import (
	todo "github.com/alexeipyp/todo_rest_api"
	"github.com/alexeipyp/todo_rest_api/pkg/repository"
	"github.com/alexeipyp/todo_rest_api/pkg/service/auth"
	todoitem "github.com/alexeipyp/todo_rest_api/pkg/service/todo_item"
	todolist "github.com/alexeipyp/todo_rest_api/pkg/service/todo_list"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, input todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoItem:      todoitem.New(repos.TodoItem, repos.TodoList),
	}
}
