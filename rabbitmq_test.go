package core

import (
	"os"
	"testing"
)

func TestRabbitMQInitialization(t *testing.T) {
	// ToDo add a server for tests and set credentials in the code below
	//// Preserving the original values of environment variables
	//originalHost := os.Getenv("RABBITMQ_HOST")
	//originalPort := os.Getenv("RABBITMQ_PORT")
	//originalUsername := os.Getenv("RABBITMQ_USERNAME")
	//originalPassword := os.Getenv("RABBITMQ_PASSWORD")
	//
	//// Setting test values of environment variables
	//os.Setenv("RABBITMQ_HOST", "test_host")
	//os.Setenv("RABBITMQ_PORT", "test_port")
	//os.Setenv("RABBITMQ_USERNAME", "test_username")
	//os.Setenv("RABBITMQ_PASSWORD", "test_password")
	//
	//defer func() {
	//	// Restoring the original values of environment variables after the test
	//	os.Setenv("RABBITMQ_HOST", originalHost)
	//	os.Setenv("RABBITMQ_PORT", originalPort)
	//	os.Setenv("RABBITMQ_USERNAME", originalUsername)
	//	os.Setenv("RABBITMQ_PASSWORD", originalPassword)
	//}()

	// Restoring the original values of environment variables after the test
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	username := os.Getenv("RABBITMQ_USERNAME")
	password := os.Getenv("RABBITMQ_PASSWORD")

	// Getting the values of environment variables for initialization
	rabbitMQClient, err := NewRabbitMQ(host, port, username, password)
	if err != nil {
		t.Errorf("Error while initializing RabbitMQ: %v", err)
	}

	// Checking that the RabbitMQ client was successfully created
	if rabbitMQClient == nil {
		t.Error("RabbitMQ client not created")
	}
}
