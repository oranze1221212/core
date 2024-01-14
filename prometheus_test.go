package core

import (
	"os"
	"testing"
)

func TestPrometheusInitialization(t *testing.T) {
	// ToDo add a server for tests and set credentials in the code below
	//// Preserving the original values of environment variables
	//originalHost := os.Getenv("PROMETHEUS_HOST")
	//originalPort := os.Getenv("PROMETHEUS_PORT")
	//
	//// Setting test values of environment variables
	//os.Setenv("PROMETHEUS_HOST", "test_host")
	//os.Setenv("PROMETHEUS_PORT", "test_port")
	//
	//defer func() {
	//	// Restoring the original values of environment variables after the test
	//	os.Setenv("PROMETHEUS_HOST", originalHost)
	//	os.Setenv("PROMETHEUS_PORT", originalPort)
	//}()

	// Getting the values of environment variables for initialization
	host := os.Getenv("PROMETHEUS_HOST")
	port := os.Getenv("PROMETHEUS_PORT")

	// Initializing the Prometheus client with parameters
	prometheusClient, err := NewPrometheus(host, port)
	if err != nil {
		t.Errorf("Error initializing Prometheus: %v", err)
	}

	// Verifying that the Prometheus client was successfully created
	if prometheusClient == nil {
		t.Error("Prometheus client not created")
	}
}
