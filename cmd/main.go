package main

import (
	"log"

	todo "github.com/alexeipyp/todo_rest_api"
	"github.com/alexeipyp/todo_rest_api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
