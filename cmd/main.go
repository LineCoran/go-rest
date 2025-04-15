package main

import (
	"log"
	"os"

	"fmt"

	"github.com/LineCoran/go-api/cmd/handler"
	"github.com/LineCoran/go-api/cmd/repository"
	"github.com/LineCoran/go-api/cmd/service"
	todo "github.com/LineCoran/go-api/pkg"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalln("Ошибка чтения файла конфигурации", err.Error())
	}

	if err := initEnv(); err != nil {
		log.Fatalln("Ошибка загрузки .env файла", err.Error())
	}

	// dbpassword := os.Getenv("DB_PASSWORD")

	db, err := repository.NewPostgresDB(repository.Config{
		Username: viper.GetString("db_username"),
		DBName:   viper.GetString("db_name"),
		Host:     viper.GetString("db_host"),
		Port:     viper.GetString("db_port"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db_sslmode"),
	})
	if err != nil {
		log.Fatalln("Ошибка создания БД", err.Error())
	}
	fmt.Println(db)
	port := viper.GetString("port")
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)

	fmt.Println("Сервер запущен на порте: ", port)
	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalln("Ошибка сервера", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initEnv() error {
	return godotenv.Load()
}
