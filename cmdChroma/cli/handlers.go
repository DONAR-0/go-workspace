package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/urfave/cli/v3"
)

// Test Connection
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

// Current tenants
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

// List databases
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

// handleListCollection List Collection
func handleListCollection(_ context.Context, cmd *cli.Command) error {
	slog.Info("List All Collections:", "Tenant", cmd.String("tenant"), "Database", cmd.String("database"))

	// Create Client
	chromaClient, err := createChromaClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Chroma client: %w", err)
	}
	// Call the method we added to your client package
	collections, err := chromaClient.ListCollections()
	if err != nil {
		slog.Error("Failed to list Databases", "error", err)
		return fmt.Errorf("could not list Databases: %w", err)
	}
	slog.Info("✅ Successfully Retrieved Collections for database: " + cmd.String("database"))

	for _, collection := range collections {
		// slog.Info("- ", "ID", db.Id, "Tenant", db.Tenant, "Name", db.Name)
		slog.Info("-", "collection", collection.Name, "ID", collection.ID)
	}

	return nil
}

// handleCreateCollection Create Collection
func handleCreateCollection(_ context.Context, cmd *cli.Command) error {
	// Get Positional Argument
	collectionName := cmd.Args().Get(0)

	// Validate if name is provided
	if collectionName == "" {
		return fmt.Errorf("collectionName name is required as the first argument")
	}

	slog.Info("Creating Collection", "name", collectionName, "total_args", cmd.Args().Len())

	// Create Client
	chromaClient, err := createChromaClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Chroma client: %w", err)
	}

	id, err := chromaClient.CreateCollection(collectionName)
	if err != nil {
		return err
	}

	slog.Info("✅ Collection created", "Name ", collectionName, "Id ", id)

	return nil
}

// List Documents in collection
func handleListDocuments(_ context.Context, c *cli.Command) error {
	slog.Info("Inside handleListDocuments", "Tenant", c.String("tenant"), "Database", c.String("database"))
	input := c.Args().Get(0)
	if input == "" {
		return fmt.Errorf("Argument empty Collection name not found, Please provide collection name")
	}
	client, _ := createChromaClient(c)

	// Step 1: Always try to resolve the name to an ID first
	targetID, err := client.GetIDByName(input)
	if err != nil {
		targetID = input
	}
	slog.Info("Resolved Collection Name to ID:", "Name", input, "ID", targetID)

	// Step 2: Now call the endpoint with a guaranteed UUID (hopefully)
	docs, err := client.ListDocuments(targetID)
	if err != nil {
		return fmt.Errorf("failed to list documents: %w", err)
	}

	slog.Info(fmt.Sprintf("✅ Retrieved %d documents from %s\n", len(docs.IDs), input))
	return nil
}

func handleAddRecordDocumentInCollection(_ context.Context, c *cli.Command) error {
	collectionName := c.Args().Get(0)
	if collectionName == "" {
		return fmt.Errorf("Argument empty collection name not found, Please propvide collection name")
	}
	client, err := createChromaClient(c)
	if err != nil {
		return err
	}
	content := c.String("doc")
	docId := c.String("id")
	if docId == "" {
		// Fallback  to simple timestamp ID if none existed
		docId = fmt.Sprintf("doc-%d", time.Now().UnixNano())
	}

	//1. Resolve Name -> UUID
	targetID, err := client.GetIDByName(collectionName)
	if err != nil {
		return err
	}
	//2. Add the document
	err = client.AddDocument(targetID, []string{docId}, []string{content}, nil)
	return nil
}
