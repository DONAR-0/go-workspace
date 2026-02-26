package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/donar0/favecli/onnx"
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
		return fmt.Errorf("argument empty Collection name not found, Please provide collection name")
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

	fmt.Printf("\n--- Documents in %s ---\n", targetID)
	for i := 0; i < len(docs.IDs); i++ {
		fmt.Printf("ID:       %s\n", docs.IDs[i])

		// Check if Documents slice isn't empty
		if len(docs.Documents) > i {
			fmt.Printf("Content:  %s\n", docs.Documents[i])
		}

		// Check if Metadata exists for this record
		if len(docs.Metadatas) > i && docs.Metadatas[i] != nil {
			fmt.Printf("Metadata: %v\n", docs.Metadatas[i])
		}
		fmt.Println("-----------------------")
	}

	slog.Info(fmt.Sprintf("✅ Retrieved %d documents from %s\n", len(docs.IDs), input))
	return nil
}

func handleAddRecordDocumentInCollection(_ context.Context, c *cli.Command) error {
	collectionName := c.Args().Get(0)
	if collectionName == "" {
		return fmt.Errorf("argument empty collection name not found, Please propvide collection name")
	}
	// 1. Create the standard HTTP Client
	client, err := createChromaClient(c)
	if err != nil {
		return err
	}

	// 2. Load the AI embedding engine (only for this command)
	slog.Info("Loading AI Embedding Engine...")
	embedder, err := onnx.NewEmbedder(
		"./models/minilm/model.onnx",
		"./models/minilm/tokenizer.json",
		"./models/onnx_runtime/onnxruntime-linux-x64-1.24.2/lib/libonnxruntime.so",
	)
	if err != nil {
		return fmt.Errorf("failed to initialize AI engine: %w", err)
	}
	defer embedder.Close() // Crucial: Frees RAM/C++ memory when done
	client.Embedder = embedder

	// 3. Prepare data
	content := c.String("doc")
	docId := c.String("id")
	if docId == "" {
		docId = fmt.Sprintf("doc-%d", time.Now().UnixNano())
	}

	// 4. Generate the Embedding Vector
	slog.Info("Generating local embedding...", "text_length", len(content))
	vector, err := client.GenerateLocalEmbedding(content)
	if err != nil {
		return fmt.Errorf("embedding failed: %w", err)
	}

	// 5. Resolve Collection Name to UUID
	targetID, err := client.GetIDByName(collectionName)
	if err != nil {
		return fmt.Errorf("could not find collection '%s': %w", collectionName, err)
	}

	// 6. Upload to ChromaDB
	slog.Info("Uploading to ChromaDB", "collection", collectionName, "id", docId)
	err = client.AddDocument(targetID, docId, content, vector)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Successfully added document '%s' to collection '%s'\n", docId, collectionName)
	return nil
}
