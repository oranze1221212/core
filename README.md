# Core Connectors Module

## Overview

The Connectors Module is a comprehensive library written in Go, designed to facilitate seamless integration with various external services like Kafka, RabbitMQ, Prometheus, Redis, and CockroachDB. It simplifies the process of establishing and managing connections, making it ideal for applications requiring reliable communication with these services.

## Features

- Easy-to-use interfaces for popular services: Kafka, RabbitMQ, Prometheus, Redis, and SQL DB servers (Postgresql, MariaDB, Mysql, CockroachDB and etc).
- Automated initialization of connectors based on environment variables.
- Thread-safe handling of service connections.
- Extensibility to add more service connectors as per future requirements.

## Installation

To incorporate this module into your Go application, import it as shown below:

```go
import (
"https://github.com/oranze1221212/core"
)
```

## Usage

### Initialization of Connectors

The module's \`init()\` function automatically initializes connectors, leveraging the following structure to manage them:

```go
type Core struct {
Mu               sync.RWMutex
KafkaProducer    sarama.SyncProducer
KafkaConsumer    sarama.Consumer
RabbitMQClient   *amqp.Connection
PrometheusClient prometheus.Client
RedisClient      *redis.Client
RedisRedjet      *redjet.Client
SQLDB            *sqlx.DB
}
```

Create an instance of the \`Core\` structure to access these connectors:

```go
core := core.Handler{}
```

### Individual Connector Usage

#### Kafka

Initialize the Kafka client:

```go
kafkaClient, err := core.NewKafka("host", "port", "username", "password")
if err != nil {
// Error handling
}
// Interaction with Kafka
```

#### RabbitMQ

Initialize the RabbitMQ client:

```go
rabbitMQClient, err := core.NewRabbitMQ("host", "port", "username", "password")
if err != nil {
// Error handling
}
// Interaction with RabbitMQ
```

#### Prometheus

Initialize the Prometheus client:

```go
prometheusClient, err := core.NewPrometheus("host", "port")
if err != nil {
// Error handling
}
// Interaction with Prometheus
```

#### Redis

Initialize the Redis client:

```go
redisClient, err := core.NewRedis("host", "port", "password")
if err != nil {
// Error handling
}
// Interaction with Redis
```

#### CockroachDB

Initialize the CockroachDB client:

```go
cockroachDBClient, err := core.NewCockroachDB("host", "port", "username", "password")
if err != nil {
// Error handling
}
// Interaction with CockroachDB
```

## Configuration

Set the following environment variables for each connector's configuration:

- **Kafka Connector:**
  - `KAFKA_ENABLED`: Kafka enable flag
  - `KAFKA_HOST`: Kafka broker host
  - `KAFKA_PORT`: Kafka broker port
  - `KAFKA_USERNAME`: Kafka username (optional)
  - `KAFKA_PASSWORD`: Kafka password (optional)

- **RabbitMQ Connector:**
  - `RABBITMQ_ENABLED`: RabbitMQ enable flag
  - `RABBITMQ_HOST`: RabbitMQ host
  - `RABBITMQ_PORT`: RabbitMQ port
  - `RABBITMQ_USERNAME`: RabbitMQ username (optional)
  - `RABBITMQ_PASSWORD`: RabbitMQ password (optional)

- **Prometheus Connector:**
  - `PROMETHEUS_ENABLED`: Prometheus enable flag
  - `PROMETHEUS_HOST`: Prometheus host
  - `PROMETHEUS_PORT`: Prometheus port

- **Redis Connector:**
  - `REDIS_ENABLED`: Redis enable flag
  - `REDIS_HOST`: Redis host
  - `REDIS_PORT`: Redis port
  - `REDIS_PASSWORD`: Redis password (optional)

- **Redis-Redjet Connector:**
  - `REDIS_REDJET_ENABLED`: Redis enable flag
  - `REDIS_REDJET_ADDR`: Redis host
  - `REDIS_REDJET_USERNAME`: Redis username
  - `REDIS_REDJET_PASSWORD`: Redis password (optional)

- **CockroachDB Connector:**
  - `SQLDB_ENABLED`: SQLDB enable flag
  - `SQLDB_HOST`: SQLDB host
  - `SQLDB_PORT`: SQLDB port
  - `SQLDB_USERNAME`: SQLDB username (optional)
  - `SQLDB_PASSWORD`: SQLDB password (optional)

## Testing

Run unit tests across the module using:

```plaintext
go test ./...
```

## Contributing

Contributions are welcome. Please adhere to the project's code of conduct and contribution guidelines when submitting patches and additions.

## License

This project is licensed under Apache License Version 2.0, see the LICENSE file for more details.
