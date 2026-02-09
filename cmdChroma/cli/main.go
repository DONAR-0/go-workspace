package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	cClient "github.com/donar0/favecli/client"
	"github.com/urfave/cli/v3"
)

const (
	// App Version Of CLI
	AppVersion  = "1.0.0"
	ExitSuccess = 1
	ExitError   = 1
)

func main() {
	// Initialize Logger first
	InitLogger()

	// Recover from panic gracefully
	defer func() {
		if r := recover(); r != nil {
			slog.Error("CLI application Panicked", "panic", r)
			fmt.Printf("Error: An Unexpected error occurred: %v\n", r)
			os.Exit(ExitError)
		}
	}()

	app := createApp()

	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Error("CLI execution failed", "error", err)
		fmt.Printf("Error: %v\n", err)
		os.Exit(ExitError)
	}
}

// createChromaClient creates a Chroma client based on CLI context
func createChromaClient(c *cli.Command) (*cClient.ChromaClient, error) {
	// For now, use the default client
	// In the future, this could be enhanced to use host/port from flags
	return cClient.NewChromaDBClient(fmt.Sprintf("http://%s:%s", c.String("host"), c.String("port")), "default_tenant", "default_database"), nil
}
