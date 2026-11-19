package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/kosatnkn/catalyst/v3/infra"
)

type Server struct {
	cfg infra.RESTConfig
	ctr *infra.Container
	srv *http.Server
	// errs delivers an unrecoverable serve error (ex: listener failure) that
	// happens after Start has returned. Request failures never reach here.
	errs chan error
}

// NewServer returns a new REST server instance.
func NewServer(cfg infra.RESTConfig, ctr *infra.Container) (*Server, error) {
	s := &Server{
		cfg: cfg,
		ctr: ctr,
	}

	return s, nil
}

// Start starts the REST server.
func (s *Server) Start() error {
	ctx := context.Background()

	s.ctr.Logger.Info(ctx, "starting REST server")
	s.ctr.Logger.Debug(ctx, fmt.Sprintf("config values: %+v", s.cfg))

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.Port),
		Handler: newRouter(s.cfg, s.ctr).Handler(),
	}

	// bind eagerly so a bind failure (ex: port already in use) is returned from Start
	// deterministically, instead of racing a timer against the serve goroutine
	ln, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return fmt.Errorf("rest, Start: failed to bind to %s, %w", s.srv.Addr, err)
	}

	// channel to catch unrecoverable serve failures after Start has returned
	s.errs = make(chan error, 1)

	go func() {
		// Serve owns ln and closes it when it returns, Stop() causes ErrServerClosed
		if err := s.srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			s.errs <- fmt.Errorf("rest, Serve: %w", err)
		}
	}()

	return nil
}

// Err returns a channel that receives an unrecoverable error if the server
// stops serving after a successful Start.
func (s *Server) Err() <-chan error {
	return s.errs
}

// Stop gracefully stops the REST server, waiting for in-flight requests until ctx is done
// and then forcefully closing any remaining connections.
func (s *Server) Stop(ctx context.Context) error {
	s.ctr.Logger.Info(context.Background(), "stopping REST server")

	if err := s.srv.Shutdown(ctx); err != nil {
		// when timeout exceeds forcefully close any remaining connections
		if cerr := s.srv.Close(); cerr != nil {
			return fmt.Errorf("rest, Stop: error in forceful connection termination, %w, %w", cerr, err)
		}

		return fmt.Errorf("rest, Stop: error in server shutdown, %w", err)
	}

	return nil
}
