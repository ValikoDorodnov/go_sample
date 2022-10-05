package http

import (
	"context"
	"net/http"
	"time"

	"github.com/ValikoDorodnov/go_sample/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewRestServer(conf config.RestConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + conf.Port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
