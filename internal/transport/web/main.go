package web

import (
	"context"
	"gostack/internal/repo/sqlite"
	"gostack/internal/services/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WebHandler struct {
	*user.UserService
}

func Serve() {
	repo := sqlite.New()
	userService := user.New(repo)
	wh := &WebHandler{UserService: userService}
	mux := wh.Register()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal to close app
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on port 4000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	<-osSignal
	log.Println("Shutdown signal received. Shutting down...")

	// Timeout to close handlers
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
	defer shutdownCancel()

	// Stop the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}
