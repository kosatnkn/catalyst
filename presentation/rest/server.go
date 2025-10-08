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

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.ctr.Logger.Error(context.Background(), errors.Join(errors.New("rest: error in server"), err).Error())
			return
		}
	}()

	return nil
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
		return err
	}

	return nil
}
