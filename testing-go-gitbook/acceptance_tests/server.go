package acceptance_tests

import (
	"context"
	"net/http"
	"os"
	"time"
)

const defaultTimeout = 5 * time.Second

type (
	HttpServer interface {
		ListenAndServe() error
		Shutdown(ctx context.Context) error
	}

	Server struct {
		shutdown <-chan os.Signal
		delegate HttpServer
		timeout  time.Duration
	}

	ServerOption func(server *Server)
)

func WithTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func WithShutdownSignal(shutdown <-chan os.Signal) ServerOption {
	return func(server *Server) {
		server.shutdown = shutdown
	}
}

func NewServer(server HttpServer, opts ...ServerOption) *Server {
	s := &Server{
		shutdown: newInterruptSignalChannel(),
		delegate: server,
		timeout:  defaultTimeout,
	}

	for _, o := range opts {
		o(s)
	}

	return s
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	select {
	case err := <-s.delegateListenAndServe():
		return err
	case <-ctx.Done():
		return s.shutdownDelegate(ctx)
	case <-s.shutdown:
		return s.shutdownDelegate(ctx)
	}
}

func (s *Server) delegateListenAndServe() chan error {
	errCh := make(chan error)

	go func() {
		if err := s.delegate.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	return errCh
}

func (s *Server) shutdownDelegate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.delegate.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return err
	}

	return ctx.Err()
}
