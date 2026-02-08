package cClient

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type ChromaClient struct {
	URL, Tenant, Database string
	client                *http.Client
}

func NewChromaDBClient(url, tenant, database string) *ChromaClient {
	slog.Info("Initiating ChromaClient Client", "URL:", url, "Tenant:", tenant, "Database:", database)
	return &ChromaClient{
		URL:      url,
		Tenant:   tenant,
		Database: database,
		client:   &http.Client{},
	}
}

func (c *ChromaClient) TestConnection() error {

	resp, err := c.client.Get(fmt.Sprintf("%s/api/v2/heartbeat", c.URL))
	if err != nil {
		return fmt.Errorf("failed to connect to ChromaDB at %s: %w", c.URL, err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("heartbeat failed with status: %d, response: %s", resp.StatusCode, string(body))
	}
	slog.Info(fmt.Sprintf("ChromaDB connection successful: %s\n", string(body)))
	return nil
}

func (c *ChromaClient) GetTenant(name string) (bool, error) {
	// Correct endpoint for checking a specific tenant
	endpoint := fmt.Sprintf("%s/api/v2/tenants/%s", c.URL, name)

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 200 means exists, 404 means it doesn't
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, fmt.Errorf("unexpected status: %d", resp.StatusCode)
}

type Database struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Tenant string `json:"tenant"`
}

func (c *ChromaClient) ListDatabases() ([]Database, error) {
	// URL includes the specific tenant from your client struct
	endpoint := fmt.Sprintf("%s/api/v2/tenants/%s/databases", c.URL, c.Tenant)
	slog.Info("List Databases from URL :" + endpoint)
	resp, err := c.client.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list databases: status %d, body: %s", resp.StatusCode, string(body))
	}

	// Chroma returns a list of database names as strings
	var databases []Database
	if err := json.NewDecoder(resp.Body).Decode(&databases); err != nil {
		return nil, fmt.Errorf("failed to decode databases: %w", err)
	}

	return databases, nil
}
