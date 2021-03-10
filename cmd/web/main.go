package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/stackbreak/tasx/internal/app/web"
	"github.com/stackbreak/tasx/internal/pkg/handlers"
	"github.com/stackbreak/tasx/internal/pkg/repositories/postgres"
	"github.com/stackbreak/tasx/internal/pkg/services"
	"github.com/stackbreak/tasx/internal/pkg/tokens"
)

func main() {
	log := web.NewLogger()
	config := web.NewConfig()

	if err := config.LoadFile(); err != nil {
		log.Fatal("error initializing config file: ", err)
	}

	if err := config.LoadEnv(); err != nil {
		log.Fatal("error initializing env variables: ", err)
	}

	db, err := postgres.NewDB(&postgres.PgConfig{
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

	tokenManager, err := tokens.NewTokenMangerHS([]byte(config.Env.TokenSecret))
	if err != nil {
		log.Fatal("error initializing token manager:", err)
	}

	repos := postgres.NewRepositoryManager(db)
	services := services.NewServices(repos, tokenManager)
	globalHandler := handlers.NewGlobalHandler(services, log)

	port := config.File.GetString("api.port")
	srv := new(web.Server)

	go func() {
		if err := srv.Run(port, globalHandler.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatal("error occurred while running server:", err)
		}
	}()

	log.Println("App started on port:", port)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	log.Println("App shutting down, received signal:", sig)

	ctxShutdown, _ := context.WithTimeout(context.Background(), 15*time.Second)
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatal("error occurred on server shutdown:", err)
	}
}
