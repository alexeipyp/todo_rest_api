package repository

import (
	todo "github.com/alexeipyp/todo_rest_api"
	postgres_auth "github.com/alexeipyp/todo_rest_api/pkg/repository/postgres/auth"
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
}

type TodoItem interface {
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
	}
}
