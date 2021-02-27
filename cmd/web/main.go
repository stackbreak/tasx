package main

import (
	"log"

	"github.com/stackbreak/tasx/internal/app/web"
	"github.com/stackbreak/tasx/internal/pkg/handlers"
	"github.com/stackbreak/tasx/internal/pkg/repository"
	"github.com/stackbreak/tasx/internal/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	globalHandler := handlers.NewGlobalHandler(services)

	srv := new(web.Server)
	if err := srv.Run("4000", globalHandler.InitRoutes()); err != nil {
		log.Fatal("Stop server because", err)
	}
}
