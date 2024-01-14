package core

import (
	"fmt"
	"os"
)

func init() {
	Handler.Mu.Lock()
	defer Handler.Mu.Unlock()

	// Initializing Prometheus
	if os.Getenv("PROMETHEUS_ENABLED") != "false" {
		host := os.Getenv("PROMETHEUS_HOST")
		port := os.Getenv("PROMETHEUS_PORT")
		client, err := NewPrometheus(host, port)
		if err != nil {
			fmt.Println("Prometheus initialization error:", err)
			os.Exit(1)
		} else {
			Handler.Prometheus = client
			Handler.PrometheusIsInitialized = true
		}
	}

	// Initializing SqlDB
	if os.Getenv("SQLDB_ENABLED") != "false" {
		host := os.Getenv("SQLDB_HOST")
		port := os.Getenv("SQLDB_PORT")
		user := os.Getenv("SQLDB_USER")
		password := os.Getenv("SQLDB_PASSWORD")
		dbName := os.Getenv("SQLDB_DB")
		dbDriver := os.Getenv("SQLDB_DRIVER")
		db, err := NewSqlDB(host, port, user, password, dbName, dbDriver)
		if err != nil {
			fmt.Println("Error connecting to SqlDB:", err)
			os.Exit(1)
		} else {
			Handler.SQLDB = db
			Handler.SQLDBIsInitialized = true
		}
	}

	// Initializing Redis
	if os.Getenv("REDIS_ENABLED") != "false" {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		password := os.Getenv("REDIS_PASSWORD")
		client, err := NewRedis(host, port, password)
		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
			os.Exit(1)
		} else {
			Handler.Redis = client
			Handler.RedisIsInitialized = true
		}
	}

	// Initializing Redis-Redjet
	if os.Getenv("REDIS_REDJET_ENABLED") != "false" {
		addr := os.Getenv("REDIS_REDJET_ADDR")
		username := os.Getenv("REDIS_REDJET_USERNAME")
		password := os.Getenv("REDIS_REDJET_PASSWORD")

		redisRedjetClient, err := NewRedisRedjet(addr, username, password)
		if err != nil {
			fmt.Printf("Error connecting to Redis-Redjet: %s\n", err)
			os.Exit(1)
		}

		Handler.RedisRedjet = redisRedjetClient
	}

	// Initializing RabbitMQ
	if os.Getenv("RABBITMQ_ENABLED") != "false" {
		host := os.Getenv("RABBITMQ_HOST")
		port := os.Getenv("RABBITMQ_PORT")
		user := os.Getenv("RABBITMQ_USER")
		password := os.Getenv("RABBITMQ_PASSWORD")
		conn, err := NewRabbitMQ(host, port, user, password)
		if err != nil {
			fmt.Println("Error connecting to RabbitMQ:", err)
			os.Exit(1)
		} else {
			Handler.RabbitMQ = conn
			Handler.RabbitMQIsInitialized = true
		}
	}

	// Initializing Kafka
	if os.Getenv("KAFKA_ENABLED") != "false" {
		host := os.Getenv("KAFKA_HOST")
		port := os.Getenv("KAFKA_PORT")
		username := os.Getenv("KAFKA_USERNAME")
		password := os.Getenv("KAFKA_PASSWORD")

		producer, consumer, err := NewKafka(host, port, username, password)
		if err != nil {
			fmt.Println("Error connecting to Kafka:", err)
			os.Exit(1)
		}

		Handler.KafkaProducer = producer
		Handler.KafkaConsumer = consumer
		Handler.KafkaIsInitialized = true
	}
}
