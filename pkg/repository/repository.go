package repository

import (
	todo "github.com/alexeipyp/todo_rest_api"
	postgres_auth "github.com/alexeipyp/todo_rest_api/pkg/repository/postgres/auth"
	postgres_todoitem "github.com/alexeipyp/todo_rest_api/pkg/repository/postgres/todo_item"
	postgres_todolist "github.com/alexeipyp/todo_rest_api/pkg/repository/postgres/todo_list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, input todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres_auth.New(db),
		TodoList:      postgres_todolist.New(db),
		TodoItem:      postgres_todoitem.New(db),
	}
}
