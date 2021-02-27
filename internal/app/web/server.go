package web

import (
	"context"
	"net/http"
	"time"
)

var MaxHeaderBytes = 1 << 20 // 1MB

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, router http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: MaxHeaderBytes,
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
