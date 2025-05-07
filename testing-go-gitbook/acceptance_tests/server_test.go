package acceptance_tests_test

import (
	"acceptance_tests"
	"acceptance_tests/assert"
	"context"
	"errors"
	"os"
	"testing"
	"time"
)

type ServerSpy struct {
	ListenAndServeFunc func() error
	listened           chan struct{}

	ShutdownFunc func() error
	shutdown     chan struct{}
}

func (s *ServerSpy) ListenAndServe() error {
	s.listened <- struct{}{}
	return s.ListenAndServeFunc()
}

func (s *ServerSpy) Shutdown(ctx context.Context) error {
	s.shutdown <- struct{}{}
	return s.ShutdownFunc()
}

func (s *ServerSpy) AssertListened(t *testing.T) {
	t.Helper()
	assert.SignalSent(t, s.listened, "listen")
}

func (s *ServerSpy) AssertShutdown(t *testing.T) {
	t.Helper()
	assert.SignalSent(t, s.shutdown, "shutdown")
}

func NewServerSpy() *ServerSpy {
	return &ServerSpy{
		listened: make(chan struct{}, 1),
		shutdown: make(chan struct{}, 1),
	}
}

func TestGracefulShutdownServer_Listen(t *testing.T) {
	t.Run("happy path, listen, wait for interrupt, shutdown gracefully", func(t *testing.T) {
		var (
			interrupt = make(chan os.Signal)
			serverSpy = NewServerSpy()
			server    = acceptance_tests.NewServer(serverSpy, acceptance_tests.WithShutdownSignal(interrupt))
			ctx       = context.Background()
		)

		serverSpy.ListenAndServeFunc = func() error {
			return nil
		}

		serverSpy.ShutdownFunc = func() error {
			return nil
		}

		go func() {
			if err := server.ListenAndServe(ctx); err != nil {
				t.Error(err)
			}
		}()

		// verify we can listen on the delegate server
		serverSpy.AssertListened(t)

		// verify we call shutdown on the delegate when an interrupt is sent
		interrupt <- os.Interrupt
		serverSpy.AssertShutdown(t)
	})

	t.Run("when listen fails, return error", func(t *testing.T) {
		var (
			interrupt = make(chan os.Signal)
			serverSpy = NewServerSpy()
			server    = acceptance_tests.NewServer(serverSpy, acceptance_tests.WithShutdownSignal(interrupt))
			err       = errors.New("listen failed")
			ctx       = context.Background()
		)

		serverSpy.ListenAndServeFunc = func() error {
			return err
		}

		gotErr := server.ListenAndServe(ctx)
		assert.Equal(t, gotErr.Error(), err.Error())
	})

	t.Run("shutdown error gets propagated", func(t *testing.T) {
		var (
			interrupt = make(chan os.Signal)
			serverSpy = NewServerSpy()
			server    = acceptance_tests.NewServer(serverSpy, acceptance_tests.WithShutdownSignal(interrupt))
			err       = errors.New("shutdown propagated")
			ctx       = context.Background()
			errChan   = make(chan error)
		)

		serverSpy.ListenAndServeFunc = func() error {
			return nil
		}

		serverSpy.ShutdownFunc = func() error {
			return err
		}

		go func() {
			errChan <- server.ListenAndServe(ctx)
		}()

		interrupt <- os.Interrupt

		select {
		case gotErr := <-errChan:
			assert.Equal(t, gotErr.Error(), err.Error())
		case <-time.After(500 * time.Millisecond):
			t.Error("timed out waiting for shutdown error to be propagated")
		}
	})

	t.Run("context passed in can trigger shutdown too", func(t *testing.T) {
		var (
			serverSpy   = NewServerSpy()
			server      = acceptance_tests.NewServer(serverSpy)
			ctx, cancel = context.WithCancel(context.Background())
		)

		serverSpy.ListenAndServeFunc = func() error {
			return nil
		}

		serverSpy.ShutdownFunc = func() error {
			return nil
		}

		go func() {
			if err := server.ListenAndServe(ctx); err != nil && err != context.Canceled {
				t.Error(err)
			}
		}()

		// assert that we can listen on the server
		serverSpy.AssertListened(t)

		cancel()
		// assert that server was shutdown
		serverSpy.AssertShutdown(t)
	})
}
