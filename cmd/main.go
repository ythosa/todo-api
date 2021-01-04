package main

import (
	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running the server: %s ", err)
	}
}
