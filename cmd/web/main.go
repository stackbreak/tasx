package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/stackbreak/tasx/internal/app/web"
	"github.com/stackbreak/tasx/internal/pkg/handlers"
	"github.com/stackbreak/tasx/internal/pkg/repository"
	"github.com/stackbreak/tasx/internal/pkg/service"
)

func main() {
	config := web.NewConfig()
	if err := config.LoadFile(); err != nil {
		log.Fatal("error initializing config file: ", err)
	}

	if err := config.LoadEnv(); err != nil {
		log.Fatal("error initializing env variables: ", err)
	}

	db, err := repository.NewPgDB(&repository.PgConfig{
		Host:    config.Env.DbHost,
		Port:    config.Env.DbPort,
		User:    config.Env.DbUser,
		Pass:    config.Env.DbPass,
		DBName:  config.Env.DbName,
		SSLMode: config.Env.DbSslMode,
	})
	if err != nil {
		log.Fatal("error initializing db:", err)
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	globalHandler := handlers.NewGlobalHandler(services)

	port := config.File.GetString("api.port")

	srv := new(web.Server)
	fmt.Println("Server start on port:", port)
	if err := srv.Run(port, globalHandler.InitRoutes()); err != nil {
		log.Fatal("error occurred while running server:", err)
	}
}
