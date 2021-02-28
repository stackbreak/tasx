package main

import (
	"log"

	"github.com/stackbreak/tasx/internal/app/web"
	"github.com/stackbreak/tasx/internal/pkg/handlers"
	"github.com/stackbreak/tasx/internal/pkg/repository"
	"github.com/stackbreak/tasx/internal/pkg/service"
)

func main() {
	config := web.NewConfig()

	if err := config.Init(); err != nil {
		log.Fatal("error initializing configs:", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	globalHandler := handlers.NewGlobalHandler(services)

	port := config.GetString("port")

	srv := new(web.Server)
	if err := srv.Run(port, globalHandler.InitRoutes()); err != nil {
		log.Fatal("error occurred while running server:", err)
	}
}
