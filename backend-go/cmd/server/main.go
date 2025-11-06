package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dremio/dqd/internal/handlers"
	"github.com/dremio/dqd/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

const (
	defaultPort         = "8080"
	defaultReadTimeout  = 15 * time.Second
	defaultWriteTimeout = 15 * time.Second
	defaultIdleTimeout  = 60 * time.Second
	maxUploadSize       = 10 * 1024 * 1024 // 10MB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := setupRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}

	slog.Info("starting DQD server", "port", port, "version", "0.12.3")
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.CORS)

	// Set a timeout for all requests
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// Routes
	r.Get("/api/about.json", handlers.GetAbout)

	// Profile analysis routes
	r.Post("/api/simple-profile", handlers.PostSimpleProfile)
	r.Post("/api/profile", handlers.PostProfile)
	r.Post("/api/profiles", handlers.PostProfiles)
	r.Post("/api/queriesjson", handlers.NotImplemented("POST /api/queriesjson"))
	r.Post("/api/reproduction", handlers.NotImplemented("POST /api/reproduction"))
	r.Post("/api/iostat", handlers.NotImplemented("POST /api/iostat"))
	r.Post("/api/ttop", handlers.NotImplemented("POST /api/ttop"))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			slog.Error("failed to write health check response", "error", err)
		}
	})

	return r
}
