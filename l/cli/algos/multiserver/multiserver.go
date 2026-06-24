package multiserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ServerManager manges multiple servers
type ServerManager struct {
	mu      sync.Mutex
	servers map[string]*http.Server
}

// NewServerManager creates a new manager
func NewServerManager() *ServerManager {
	return &ServerManager{
		servers: make(map[string]*http.Server),
	}
}

// helloWorldServer handler
func helloWorldServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "Hello, World")
	} else {
		_, _ = fmt.Fprintf(w, "Hello, World Supposed to be / not at %s", r.URL.Path)
	}
}

// StartServer starts a new HTTP server on a given port
func (sm *ServerManager) StartServer(port string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, exists := sm.servers[port]; exists {
		slog.Default().Warn("Server already running", "port", port)
		return
	}

	srv := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(helloWorldServer),
	}

	sm.servers[port] = srv

	go func() {
		slog.Default().Info("Starting server", "port", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Default().Error("Server failed", "port", port, "error", err)
		}
	}()
}

// StopServer gracefully stops one server
func (sm *ServerManager) StopServer(port string) {
	sm.mu.Lock()
	srv, ok := sm.servers[port]
	sm.mu.Unlock()

	if !ok {
		slog.Default().Warn("Server not found", "port", port)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Default().Error("Error shutting down server", "port", port, "error", err)
	} else {
		slog.Default().Info("Server stopped", "port", port)
	}

	sm.mu.Lock()
	delete(sm.servers, port)
	sm.mu.Unlock()
}

// StopAll stops all servers gracefully
func (sm *ServerManager) StopAll() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	for port, srv := range sm.servers {
		slog.Default().Info("Stopping server", "port", port)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = srv.Shutdown(ctx)

		cancel()
	}

	sm.servers = make(map[string]*http.Server)

	slog.Default().Info("All servers stopped")
}

// ListServers returns the ports currently running
func (sm *ServerManager) ListServers() []string {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	ports := make([]string, 0, len(sm.servers))
	for port := range sm.servers {
		ports = append(ports, port)
	}

	return ports
}

// Run starts multiple servers and waits for interrupt to stop all
func Run() {
	manager := NewServerManager()
	for i := 5000; i <= 5002; i++ {
		manager.StartServer(fmt.Sprintf(":%d", i))
	}

	// Handle Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	manager.StopAll()
}
