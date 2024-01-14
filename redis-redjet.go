package core

import (
	"context"
	"fmt"
	"github.com/coder/redjet"
	"time"
)

func NewRedisRedjet(addr, username, password string) (*redjet.Client, error) {
	client := redjet.New(addr)
	client.AuthUsername = username
	client.AuthPassword = password
	client.ConnectionPoolSize = 10       // Connection pool
	client.IdleTimeout = 5 * time.Minute // Idle timeout

	// Connection checking
	pipeline := client.Command(context.Background(), "PING")
	err := pipeline.Ok()
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis-Redjet: %w", err)
	}

	return client, nil
}
