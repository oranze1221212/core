package core

import (
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockProducer is a mock for Kafka Producer
type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	args := m.Called(msg)
	return args.Get(0).(int32), args.Get(1).(int64), args.Error(2)
}

func (m *MockProducer) Close() error {
	args := m.Called()
	return args.Error(0)
}

// TestKafkaProducerSendMessage tests sending a message through Kafka Producer
func TestKafkaProducerSendMessage(t *testing.T) {
	// Creating a mock for Producer
	mockProducer := new(MockProducer)
	message := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("Hello Kafka")}

	// Setting expectations for the mock
	mockProducer.On("SendMessage", message).Return(int32(0), int64(0), nil)
	mockProducer.On("Close").Return(nil)

	// Testing sending a message
	partition, offset, err := mockProducer.SendMessage(message)

	// Checking the result
	assert.NoError(t, err)
	assert.Equal(t, int32(0), partition)
	assert.Equal(t, int64(0), offset)

	// Checking that all expected calls have been made
	mockProducer.AssertExpectations(t)
}

// Producer Closure Test
func TestKafkaProducerClose(t *testing.T) {
	mockProducer := new(MockProducer)
	mockProducer.On("Close").Return(nil)

	// Testing Producer Closing
	err := mockProducer.Close()

	// Checking that there were no errors
	assert.NoError(t, err)

	// Checking that all expected calls have been made
	mockProducer.AssertExpectations(t)
}
