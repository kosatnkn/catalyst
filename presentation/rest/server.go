package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

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
	s.ctr.Logger.Debug(context.Background(), fmt.Sprintf("config values: %+v", s.cfg))

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
	wait, err := time.ParseDuration(s.cfg.Wait)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		// when timeout exceeds forcefully close any remaining connections
		if cerr := s.srv.Close(); cerr != nil {
			return errors.Join(err, cerr)
		}

		return err
	}

	return nil
}
