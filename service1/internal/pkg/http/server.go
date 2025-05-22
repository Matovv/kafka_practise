package http

import (
	"context"
	"fmt"
	"local/service1/internal/pkg/configs"
	"net/http"
	"time"
)

const readHeaderTimeout = 5 * time.Second

type Server struct {
	server    *http.Server
}

func NewServer(config *configs.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr: config.BindAddr,
			Handler: handler,
			ReadHeaderTimeout: readHeaderTimeout,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	fmt.Println("starting the server.")
	if err := s.server.ListenAndServe(); err != nil {
		fmt.Println("error while starting the server")
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("stopping the server.")
	err := s.server.Shutdown(ctx)
	if err != nil {
		fmt.Println("error while shutting down the server.")
		return err
	}
	return nil
}