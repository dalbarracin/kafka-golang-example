package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	Config   *ConsumerConfig
}

func (c *KafkaConsumer) Read() error {

	fmt.Printf("consuming message from topic..")

	msg, err := c.consumer.ReadMessage(-1)

	if err == nil {
		fmt.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		return fmt.Errorf("consumer error: %v (%v)", err, msg)
	}

	return nil
}

func (c *KafkaConsumer) Build() error {

	nc, err := kafka.NewConsumer(c.consumerConfigMap())

	if err != nil {
		return err
	}

	c.consumer = nc

	err = c.consumer.Subscribe(c.Config.Topic, nil)

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
		"bootstrap.servers": c.Config.BootstrapServers,
		"auto.offset.reset": c.Config.AutoOffsetReset,
		"group.id":          c.Config.GroupId,
	}
}
