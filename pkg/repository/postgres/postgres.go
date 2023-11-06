package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Users      = "users"
	TodoLists  = "todo_lists"
	UsersLists = "users_lists"
	TodoItems  = "todo_items"
	ListsItems = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func New(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
