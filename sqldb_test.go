package core

import (
	"os"
	"testing"
)

func TestSQLDBInitialization(t *testing.T) {
	// ToDo add a server for tests and set credentials in the code below
	//// add a server for tests and set credentials in the code below
	//originalHost := os.Getenv("SQLDB_HOST")
	//originalPort := os.Getenv("SQLDB_PORT")
	//originalUsername := os.Getenv("SQLDB_USERNAME")
	//originalPassword := os.Getenv("SQLDB_PASSWORD")
	//originalDatabase := os.Getenv("SQLDB_DATABASE")
	//
	//// Setting test values of environment variables
	//os.Setenv("SQLDB_HOST", "test_host")
	//os.Setenv("SQLDB_PORT", "test_port")
	//os.Setenv("SQLDB_USERNAME", "test_username")
	//os.Setenv("SQLDB_PASSWORD", "test_password")
	//os.Setenv("SQLDB_DATABASE", "test_database")
	//
	//defer func() {
	//	// Restoring the original values of environment variables after the test
	//	os.Setenv("SQLDB_HOST", originalHost)
	//	os.Setenv("SQLDB_PORT", originalPort)
	//	os.Setenv("SQLDB_USERNAME", originalUsername)
	//	os.Setenv("SQLDB_PASSWORD", originalPassword)
	//	os.Setenv("SQLDB_DATABASE", originalDatabase)
	//}()

	// Restoring the original values of environment variables after the test
	host := os.Getenv("SQLDB_HOST")
	port := os.Getenv("SQLDB_PORT")
	username := os.Getenv("SQLDB_USERNAME")
	password := os.Getenv("SQLDB_PASSWORD")
	database := os.Getenv("SQLDB_DATABASE")

	// Initialize the NewSqlDB client with parameters
	sqlDB, err := NewSqlDB(host, port, username, password, database, "postgres")
	if err != nil {
		t.Errorf("Error initializing NewSqlDB: %v", err)
	}

	// Verifying that the NewSqlDB client was successfully created
	if sqlDB == nil {
		t.Error("NewSqlDB client not created")
	}
}
