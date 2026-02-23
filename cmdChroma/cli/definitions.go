package main

import (
	"context"
	"log/slog"

	"github.com/urfave/cli/v3"
)

func createApp() *cli.Command {
	return &cli.Command{
		Name:    "chroma",
		Version: AppVersion,
		Usage:   "Command Line Inteface for Chroma DB Operations",
		Description: "A Comprehensive CLI Tool to interact with Chroma DB," +
			"including connection testing, data operations, and more",
		// Global Flags Available to all command
		Flags: []cli.Flag{hostFlag, portFlag, verboseFlag},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			if c.Bool("verbose") {
				slog.Info("Verbose mode enabled")
			}
			return ctx, nil
		},
		Commands: []*cli.Command{
			TestCommandDefinition(),
			GetTenantDefinition(),
			ListDatabaseDefinition(),
			ListCollectionsInDatabaseDefinition(),
			GetOrCreateCollectionInDatabaseDefinition(),
			ListRecordsInCollection(),
			AddDocumentCommandDefinition(),
		},

		//Default action when no command is provided
		Action: func(ctx context.Context, c *cli.Command) error {
			return cli.ShowAppHelp(c)
		},
	}
}

func GetTenantDefinition() *cli.Command {
	return &cli.Command{
		Name:    "tenant", // Better to use lowercase short names for CLI commands
		Aliases: []string{"currentTenant", "cT"},
		Usage:   "Current Tenant in Chroma DB",
		Action:  handleCurrentTenants,
		Flags:   []cli.Flag{tenantFlag},
	}
}

func ListDatabaseDefinition() *cli.Command {
	return &cli.Command{
		Name:    "databases",
		Aliases: []string{"ls-dbs", "dbs"},
		Usage:   "List Databases in current Tenant",
		Action:  handleListDatabases,
		Flags:   []cli.Flag{tenantFlag, databaseFlag},
	}
}

func ListCollectionsInDatabaseDefinition() *cli.Command {
	return &cli.Command{
		Name:    "collections",
		Aliases: []string{"ls-colls", "colls"},
		Usage:   "List All the connections in database",
		Action:  handleListCollection,
		Flags:   []cli.Flag{tenantFlag, databaseFlag},
	}
}

func GetOrCreateCollectionInDatabaseDefinition() *cli.Command {
	return &cli.Command{
		Name:    "createCollections",
		Aliases: []string{"mkdir-colls", "mkColl"},
		Usage:   "List All the connections in database",
		Action:  handleCreateCollection,
		Flags:   []cli.Flag{tenantFlag, databaseFlag},
	}
}

func ListRecordsInCollection() *cli.Command {
	return &cli.Command{
		Name:      "records",
		Aliases:   []string{"ls-rs", "rs"},
		Usage:     "List All the records in database",
		ArgsUsage: collection_args_usage,
		Action:    handleListDocuments,
		Flags:     []cli.Flag{tenantFlag, databaseFlag, collectionFlag},
	}
}

func TestCommandDefinition() *cli.Command {
	return &cli.Command{
		Name:    "testConnection",
		Aliases: []string{"test", "t"},
		Usage:   "Test the connection to Chroma DB",
		Description: "Verifies connectivity to the DB instance and " +
			"ensures the service is responding correctly",
		Action: handleTestConnection,
		Flags:  []cli.Flag{timeoutFlag, tenantFlag, databaseFlag},
	}
}

func AddDocumentCommandDefinition() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Aliases:   []string{"a", "insert"},
		Usage:     "Add a single document to collection",
		ArgsUsage: "<collection_name>",
		// Use Description to provide detailed usage examples
		Description: `Add a single document and its metadata to a collection.
		
EXAMPLES:
   # Add a document with an auto-generated ID
   chroma add my_collection --doc "this is my text"

   # Add a document with a specific ID
   chroma add my_collection --doc "another text" --id "user-001"`,
		Action: handleAddRecordDocumentInCollection,
		Flags:  []cli.Flag{docFlag, idFlag},
	}
}

// Flags
var (
	tenantFlag = &cli.StringFlag{
		Name:    "tenant",
		Value:   "default_tenant",
		Usage:   "Chroma DB Tenant",
		Sources: cli.EnvVars("TENANT"),
	}

	databaseFlag = &cli.StringFlag{
		Name:    "database",
		Value:   "default_database",
		Usage:   "Chroma DB database",
		Sources: cli.EnvVars("DATABASE"),
	}

	collectionFlag = &cli.StringFlag{
		Name:    "collection",
		Usage:   "Chroma DB collection",
		Sources: cli.EnvVars("COLLECTION"),
	}

	hostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"H"},
		Value:   "localhost",
		Usage:   "Chroma DB host address",
		Sources: cli.EnvVars("CHROMA_HOST"),
	}

	portFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "8000",
		Usage:   "Chroma DB port number",
		Sources: cli.EnvVars("CHROMA_PORT"),
	}

	verboseFlag = &cli.BoolFlag{
		Name:  "verbose",
		Usage: "Enable verbose logging",
	}

	timeoutFlag = &cli.IntFlag{
		Name:  "timeout",
		Value: 30,
		Usage: "Connection timeout in seconds",
	}

	docFlag = &cli.StringFlag{
		Name:     "doc",
		Aliases:  []string{"d"},
		Usage:    "The Text Content of the document",
		Required: true,
	}

	idFlag = &cli.StringFlag{
		Name:  "id",
		Usage: "Unique Id for the docuement (auto-generated if empty)",
		Value: "",
	}
)

const (
	collection_args_usage = "<collection_name>"
)
