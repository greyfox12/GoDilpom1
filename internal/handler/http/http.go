package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/greyfox12/GoDilpom1/internal/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg config.HTTPConfig) *Server {
	server := &Server{
		&http.Server{
			Addr: cfg.HostAddress,
			//ReadTimeout: cfg.ReadTimeout,
			//WriteTimeout: cfg.WriteTimeout,
		},
	}

	return server
}

func (s *Server) RegisterRoutes(r *Router) {
	s.srv.Handler = r.router
}

func (s *Server) Start() error {
	if s.srv.Handler == nil {
		return fmt.Errorf("no routes have registered")
	}

	err := s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}

func (s *Server) ListenServe() error {
	return s.srv.ListenAndServe()
}
