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

	// channel to catch early startup failures before the server is stable
	serr := make(chan error, 1)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serr <- err // pass err through channel to calling routine
			return
		}
		close(serr)
	}()

	// give it a moment to fail fast on startup (ex: port already in use)
	select {
	case err := <-serr:
		return err // ex: bind failed
	case <-time.After(50 * time.Millisecond):
		return nil // server is up and stable
	}
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
