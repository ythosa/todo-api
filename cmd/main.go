package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/Inexpediency/todo-rest-api"
	"github.com/Inexpediency/todo-rest-api/pkg/handler"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
	"github.com/Inexpediency/todo-rest-api/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing config: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("dbhost"),
		Port:     viper.GetString("dbport"),
		Username: viper.GetString("dbuser"),
		Password: viper.GetString("dbpass"),
		DBName:   viper.GetString("dbname"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initalize db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running the server: %s ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
