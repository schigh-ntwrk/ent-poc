package httpx

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/schigh-ntwrk/ent-poc/internal/config"
)

type Option func(*Server)

func New(opts ...Option) (*Server, error) {
	s := Server{}
	for _, f := range opts {
		f(&s)
	}

	s.done = make(chan error)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	s.stopChan = stopChan

	router := mux.NewRouter()
	for _, route := range s.routes {
		router.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	server, ok := s.server.(*http.Server)
	if !ok || server == nil {
		server = &http.Server{}
	}
	server.Handler = router
	server.Addr = fmt.Sprintf(":%d", s.port)
	s.server = server

	return &s, nil
}

func WithConfig(cfg *config.Config) Option {
	return func(s *Server) {
		s.ingressTimeout = cfg.Server.IngressTimeout
		s.shutdownTimeout = cfg.Server.ShutdownTimeout
		s.port = cfg.Server.ServicePort
	}
}

func WithHTTPServer(server *http.Server) Option {
	return func(s *Server) {
		s.server = server
	}
}

func WithRoutes(routes ...Route) Option {
	return func(s *Server) {
		s.routes = append(s.routes, routes...)
	}
}
