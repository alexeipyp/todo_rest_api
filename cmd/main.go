package main

import (
	"log"
	"os"

	todo "github.com/alexeipyp/todo_rest_api"
	"github.com/alexeipyp/todo_rest_api/pkg/handler"
	"github.com/alexeipyp/todo_rest_api/pkg/repository"
	"github.com/alexeipyp/todo_rest_api/pkg/repository/postgres"
	"github.com/alexeipyp/todo_rest_api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error initializing env variables: %s", err.Error())
	}

	db, err := postgres.New(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.New(db)
	servises := service.New(repos)
	handlers := handler.New(servises)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
