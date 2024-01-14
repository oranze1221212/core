package core

import (
	"context"
	"os"
	"testing"
)

// TestNewRedisRedjet tests the creation of a Redis-Redjet client.
func TestNewRedisRedjet(t *testing.T) {
	addr := "localhost:6379" // Specify the Redis server address for testing
	username := ""           // Specify username if necessary
	password := ""           // Specify password if necessary

	client, err := NewRedisRedjet(addr, username, password)
	if err != nil {
		t.Fatalf("Error creating Redis-Redjet client: %v", err)
	}
	defer client.Close()

	// Testing the connection using the PING command
	pipeline := client.Command(context.Background(), "PING")
	if err := pipeline.Ok(); err != nil {
		t.Errorf("Error executing PING command: %v", err)
	}
}

// Example of an integration test with Redis
func TestRedisRedjetIntegration(t *testing.T) {
	// This test requires access to a running Redis server
	if os.Getenv("RUN_INTEGRATION_TESTS") != "1" {
		t.Skip("Skipping integration tests without RUN_INTEGRATION_TESTS=1 environment variable set")
	}

	addr := "localhost:6379" // Specify the Redis server address for testing
	username := ""           // Specify username if necessary
	password := ""           // Specify password if necessary

	client, err := NewRedisRedjet(addr, username, password)
	if err != nil {
		t.Fatalf("Error creating Redis-Redjet client: %v", err)
	}
	defer client.Close()

	// Testing writing to and reading from Redis
	key, value := "testkey", "testvalue"
	setCmd := client.Command(context.Background(), "SET", key, value)
	if err := setCmd.Ok(); err != nil {
		t.Fatalf("Error setting value: %v", err)
	}

	getCmd := client.Command(context.Background(), "GET", key)
	result, err := getCmd.Bytes()
	if err != nil {
		t.Fatalf("Error getting value: %v", err)
	}
	if string(result) != value {
		t.Errorf("Expected %s, got %s", value, string(result))
	}
}
