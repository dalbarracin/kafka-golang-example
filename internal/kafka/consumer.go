package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	config   *ConsumerConfig
}

func (c *KafkaConsumer) SetConfig(config *ConsumerConfig) {
	c.config = config
}

func (c *KafkaConsumer) Read() (string, error) {

	fmt.Printf("consuming message from topic..")

	msg, err := c.consumer.ReadMessage(-1)
	if err != nil {
		return "", fmt.Errorf("consumer error: %v (%v)", err, msg)
	}

	message := string(msg.Value)

	fmt.Printf("message on %s: %s\n", msg.TopicPartition, message)

	return message, nil
}

func (c *KafkaConsumer) Build() error {

	err := c.config.Validate()
	if err != nil {
		return err
	}

	configMap := c.consumerConfigMap()

	nc, err := kafka.NewConsumer(configMap)
	if err != nil {
		return err
	}

	c.consumer = nc

	err = c.consumer.Subscribe(c.config.Topic, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *KafkaConsumer) Close() {
	c.consumer.Close()
}

func (c *KafkaConsumer) consumerConfigMap() *kafka.ConfigMap {

	return &kafka.ConfigMap{
		"bootstrap.servers": c.config.BootstrapServers,
		"auto.offset.reset": c.config.AutoOffsetReset,
		"group.id":          c.config.GroupId,
	}
}
