package httpx

import (
	"context"
	"net/http"
	"os"
	"time"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

type Server struct {
	port            int
	ingressTimeout  time.Duration
	shutdownTimeout time.Duration
	done            chan error
	stopChan        <-chan os.Signal
	server          interface {
		ListenAndServe() error
		Shutdown(context.Context) error
	}
	middleware middlewareChain
	routes     []Route
}

func (s *Server) Done() <-chan error {
	return s.done
}

func (s *Server) Start() {
	go func() {
		err := s.server.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		select {
		case s.done <- err:
		default:
		}
	}()

	<-s.stopChan
	select {
	case s.done <- nil:
	default:
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	// We allow ingress traffic for a brief moment to allow for any lag between
	// SIGTERM and HA.  This is more a Kubernetes thing, but is more generally a
	// good practice in orchestrated environments
	waitCtx, waitCancel := context.WithTimeout(ctx, s.ingressTimeout)
	defer waitCancel()
	<-waitCtx.Done()

	// We allow any in-flight requests to finish before the shutdown timeout
	// lapses. Once server.Shutdown() is called, the server accepts no new requests.
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer shutdownCancel()

	return s.server.Shutdown(shutdownCtx)
}
