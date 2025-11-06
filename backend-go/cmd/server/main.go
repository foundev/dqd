package main

import (
	"fmt"
	"log"
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

	log.Printf("Starting DQD server on port %s...", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
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

	// Placeholder routes for future implementation
	r.Post("/api/profile", handlers.NotImplemented("POST /api/profile"))
	r.Post("/api/simple-profile", handlers.NotImplemented("POST /api/simple-profile"))
	r.Post("/api/profiles", handlers.NotImplemented("POST /api/profiles"))
	r.Post("/api/queriesjson", handlers.NotImplemented("POST /api/queriesjson"))
	r.Post("/api/reproduction", handlers.NotImplemented("POST /api/reproduction"))
	r.Post("/api/iostat", handlers.NotImplemented("POST /api/iostat"))
	r.Post("/api/ttop", handlers.NotImplemented("POST /api/ttop"))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}
