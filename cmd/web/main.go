package main

import (
	"log"

	"github.com/stackbreak/tasx/internal/app/web"
)

func main() {
	srv := new(web.Server)
	if err := srv.Run("4000"); err != nil {
		log.Fatal("Stop server because", err)
	}
}
