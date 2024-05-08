package server

import (
	"fmt"
	"github.com/alishchenko/discountaria/internal/config"
	"github.com/alishchenko/discountaria/internal/data/postgres"
	"github.com/alishchenko/discountaria/internal/lib"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"os"
)

type service struct {
	// Base configs
	log *slog.Logger
	db  *postgres.DB

	// Custom configs
	cfg config.Config
}

func (s *service) run() error {
	s.log.Info(fmt.Sprintf("Service started %s", s.cfg.HTTPServer.Address))
	r := s.router()

	httpServer := s.cfg.HTTPServer
	srv := http.Server{
		Addr:         httpServer.Address,
		Handler:      r,
		ReadTimeout:  httpServer.Timeout,
		WriteTimeout: httpServer.Timeout,
		IdleTimeout:  httpServer.IdleTimeout,
	}
	return srv.ListenAndServe()
}

func newService(cfg config.Config, log *slog.Logger) *service {
	db, err := postgres.NewDB(cfg.DB.Url)
	if err != nil {
		log.Error("failed to init data", lib.Err(err))
		os.Exit(1)
	}

	return &service{
		log: log,
		db:  db,
		cfg: cfg,
	}
}
func Run(cfg config.Config, log *slog.Logger) {
	if err := newService(cfg, log).run(); err != nil {
		log.Error(errors.Wrap(err, "failed to run a service").Error())
		panic(errors.Wrap(err, "failed to run a service"))
	}
}
