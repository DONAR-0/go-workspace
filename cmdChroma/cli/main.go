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

func createApp() *cli.Command {

	return &cli.Command{
		Name:    "chroma",
		Version: AppVersion,
		Usage:   "Command Line Inteface for Chroma DB Operations",
		Description: "A Comprehensive CLI Tool to interact with Chroma DB," +
			"including connection testing, data operations, and more",
		// Global Flags Available to all command
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"H"},
				Value:   "localhost",
				Usage:   "Chroma DB host address",
				Sources: cli.EnvVars("CHROMA_HOST"),
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "8000",
				Usage:   "Chroma DB port number",
				Sources: cli.EnvVars("CHROMA_PORT"),
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose logging",
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			if c.Bool("verbose") {
				slog.Info("Verbose mode enabled")
			}
			return ctx, nil
		},
		Commands: []*cli.Command{testCommandDefinition(), GetTenantDefinition(), GetDatabaseDefinition()},

		//Default action when no command is provided
		Action: func(ctx context.Context, c *cli.Command) error {
			return cli.ShowAppHelp(c)
		},
	}
}

func handleTestConnection(ctx context.Context, cmd *cli.Command) error {
	slog.Info("Starting connection test", "host", cmd.String("host"), "port", cmd.String("port"), "timeout", cmd.Int("timeout"))

	// Create Client with context from global flags
	chromaClient, err := createChromaClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Chroma client: %w", err)
	}

	// Test the handleTestConnection
	if err := chromaClient.TestConnection(); err != nil {
		slog.Error("Connection test failed", "error", err)
		return fmt.Errorf("connection test failed: %w", err)
	}

	slog.Info("Connection test successful")
	fmt.Println("✅ Successfully connected to Chroma DB")
	return nil
}

func handleCurrentTenants(ctx context.Context, cmd *cli.Command) error {
	slog.Info("Getting Current Tenant:", "Tenant:", cmd.String("tenant"))

	// Create Client
	chromaClient, err := createChromaClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Chroma client: %w", err)
	}

	// Call the method we added to your client package
	tenantExists, err := chromaClient.GetTenant(cmd.String("tenant"))
	if err != nil {
		slog.Error("Failed to list tenants", "error", err)
		return fmt.Errorf("could not list tenants: %w", err)
	}

	slog.Info("✅ Retrieving if current tenant Exists: " + cmd.String("tenant"))
	slog.Info("Tenant Exists: " + fmt.Sprintf("%t", tenantExists))

	return nil
}

func handleListDatabases(ctx context.Context, cmd *cli.Command) error {
	slog.Info("List All Databases:", "Tenant:", cmd.String("tenant"))

	// Create Client
	chromaClient, err := createChromaClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Chroma client: %w", err)
	}

	// Call the method we added to your client package
	dbs, err := chromaClient.ListDatabases()
	if err != nil {
		slog.Error("Failed to list Databases", "error", err)
		return fmt.Errorf("could not list Databases: %w", err)
	}
	slog.Info("✅ Successfully Retrieved Databases for tenant: " + cmd.String("tenant"))

	for _, db := range dbs {
		slog.Info("- ", "ID", db.Id, "Tenant", db.Tenant, "Name", db.Name)
	}

	return nil
}

// createChromaClient creates a Chroma client based on CLI context
func createChromaClient(c *cli.Command) (*cClient.ChromaClient, error) {
	// For now, use the default client
	// In the future, this could be enhanced to use host/port from flags
	return cClient.NewChromaDBClient(fmt.Sprintf("http://%s:%s", c.String("host"), c.String("port")), "default_tenant", "default_database"), nil
}

func testCommandDefinition() *cli.Command {
	return &cli.Command{
		Name:    "testConnection",
		Aliases: []string{"test", "t"},
		Usage:   "Test the connection to Chroma DB",
		Description: "Verifies connectivity to the DB instance and " +
			"ensures the service is responding correctly",
		Action: handleTestConnection,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "timeout",
				Value: 30,
				Usage: "Connection timeout in seconds",
			},
			&cli.StringFlag{
				Name:    "tenant",
				Value:   "default_tenant",
				Usage:   "Chroma DB Tenant",
				Sources: cli.EnvVars("TENANT"),
			},
			&cli.StringFlag{
				Name:    "database",
				Value:   "default_database",
				Usage:   "Chroma DB database",
				Sources: cli.EnvVars("DATABASE"),
			},
		},
	}
}

func GetTenantDefinition() *cli.Command {
	return &cli.Command{
		Name:    "tenant", // Better to use lowercase short names for CLI commands
		Aliases: []string{"currentTenant", "cT"},
		Usage:   "Current Tenant in Chroma DB",
		Action:  handleCurrentTenants,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "tenant",
				Value:   "default_tenant",
				Usage:   "Chroma DB Tenant",
				Sources: cli.EnvVars("TENANT"),
			},
		},
	}
}

func GetDatabaseDefinition() *cli.Command {
	return &cli.Command{
		Name:    "databases",
		Aliases: []string{"ls-dbs", "dbs"},
		Usage:   "List Databases in current Tenant",
		Action:  handleListDatabases,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "tenant",
				Value:   "default_tenant",
				Usage:   "Chroma DB Tenant",
				Sources: cli.EnvVars("TENANT"),
			},
		},
	}
}
