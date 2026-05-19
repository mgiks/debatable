package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	logger *slog.Logger
}

type config struct {
	env    string
	port   string
	server serverConfig
}

type serverConfig struct {
	writeTimeout string
	readTimeout  string
	idleTimeout  string
}

func (app application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app application) run(mux http.Handler) error {
	writeTimeout, err := time.ParseDuration(app.config.server.writeTimeout)
	if err != nil {
		return fmt.Errorf("failed to parse server's write timeout: %w", err)
	}

	readTimeout, err := time.ParseDuration(app.config.server.readTimeout)
	if err != nil {
		return fmt.Errorf("failed to parse server's read timeout: %w", err)
	}

	idleTimeout, err := time.ParseDuration(app.config.server.idleTimeout)
	if err != nil {
		return fmt.Errorf("failed to parse server's idle timeout: %w", err)
	}

	srv := http.Server{
		Addr:         app.config.port,
		Handler:      mux,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}

	app.logger.Info("server has started", "port", srv.Addr, "env", app.config.env)
	return srv.ListenAndServe()
}
