package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kosatnkn/catalyst/v3/infra"
)

type Server struct {
	cfg infra.RESTConfig
	ctr *infra.Container
	srv *http.Server
	// wait is how long Stop waits for in-flight requests before forcefully closing connections
	wait time.Duration
	// errs delivers an unrecoverable serve error (ex: listener failure) that
	// happens after Start has returned. Request failures never reach here.
	errs chan error
}

// NewServer returns a new REST server instance.
func NewServer(cfg infra.RESTConfig, ctr *infra.Container) (*Server, error) {
	// parse here (not in Stop) so that a bad value fails at startup
	// instead of skipping the graceful shutdown
	wait, err := time.ParseDuration(cfg.Wait)
	if err != nil {
		return nil, fmt.Errorf("rest, NewServer: invalid wait duration %q, %w", cfg.Wait, err)
	}

	s := &Server{
		cfg:  cfg,
		ctr:  ctr,
		wait: wait,
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

	// channel to catch serve failures, both early (startup) and late (runtime)
	s.errs = make(chan error, 1)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.errs <- err // pass err through channel to calling routine
		}
	}()

	// give it a moment to fail fast on startup (ex: port already in use)
	select {
	case err := <-s.errs:
		return err // ex: bind failed
	case <-time.After(50 * time.Millisecond):
		return nil // server is up and stable
	}
}

// Err returns a channel that receives an unrecoverable error if the server
// stops serving after a successful Start.
func (s *Server) Err() <-chan error {
	return s.errs
}

// Stop stops the REST server.
func (s *Server) Stop() error {
	s.ctr.Logger.Info(context.Background(), "stopping REST server")

	ctx, cancel := context.WithTimeout(context.Background(), s.wait)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		// when timeout exceeds forcefully close any remaining connections
		if cerr := s.srv.Close(); cerr != nil {
			return fmt.Errorf("rest, Stop: error in forceful connection termination, %w, %w", cerr, err)
		}

		return fmt.Errorf("rest, Stop: error in server shutdown, %w", err)
	}

	return nil
}
