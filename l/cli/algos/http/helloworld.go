package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func helloWorldServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "Hello, World")
	} else {
		_, _ = fmt.Fprintf(w, "Hello, World Supposed to be / not at %s", r.URL.Path)
	}
}

func Serve(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(helloWorldServer),
	}

	// Channel to listen for OS signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run Server in goroutine
	go func() {
		slog.Default().Info("Starting server", "port", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Default().Error("Server error", "error", err)
		}
	}()
	//Block until we receive a stop signal
	<-stop
	slog.Default().Info("Shutting down server...")

	//Create a context with timeout to gracefully shutdown
	ctx, canel := context.WithTimeout(context.Background(), 5*time.Second)
	defer canel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Default().Error("Server forced to shutdown", "error", err)
	}
}
