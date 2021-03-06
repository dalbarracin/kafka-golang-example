package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	config   *ProducerConfig
}

func (c *KafkaProducer) SetConfig(config *ProducerConfig) {
	c.config = config
}

func (p *KafkaProducer) Write(word string) error {

	// Wait for message deliveries before shutting down
	defer p.producer.Flush(15 * 1000)

	p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.config.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(word),
	}, nil)

	events := <-p.producer.Events()

	message := events.(*kafka.Message)
	if message.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %v", message.TopicPartition)
	}

	fmt.Printf("Delivered message to %v\n", message.TopicPartition)

	return nil
}

func (p *KafkaProducer) Build() error {

	err := p.config.Validate()
	if err != nil {
		return err
	}

	configMap := p.producerConfigMap()

	np, err := kafka.NewProducer(configMap)
	if err != nil {
		return err
	}

	p.producer = np

	return nil
}

func (c *KafkaProducer) Close() {
	c.producer.Close()
}

func (p *KafkaProducer) producerConfigMap() *kafka.ConfigMap {

	return &kafka.ConfigMap{
		"bootstrap.servers": p.config.BootstrapServers,
	}
}
