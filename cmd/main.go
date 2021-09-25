package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	server "newProjectFolders/testTask1"
	"newProjectFolders/testTask1/adapters"
	"newProjectFolders/testTask1/pkg/handler"
	"newProjectFolders/testTask1/pkg/repository"
	"newProjectFolders/testTask1/pkg/service"
	"os"
)

func main() {
	config := adapters.ParseConfig()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env file: %s", err.Error())
	}
	db, err := adapters.Init(adapters.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(config.Port, handlers.Routes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
