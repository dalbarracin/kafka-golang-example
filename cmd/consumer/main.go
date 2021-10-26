package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v2"
	"github.com/dalbarracin/kafka-golang-example/internal/kafka"
)

type confg struct {
	Kafka struct {
		BootstrapServers string `conf:"default:kafka:9091"`
		AutoOffsetReset  string `conf:"default:earliest"`
		GroupId          string `conf:"default:KAFKA_CONSUMER_GROUP"`
		Topic            string `conf:"default:topic_1"`
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {

	consumerConfig, err := parseConfigValues()
	if err != nil {
		return err
	}

	consumer := &kafka.KafkaConsumer{}

	consumer.SetConfig(consumerConfig)

	err = consumer.Build()
	if err != nil {
		return err
	}

	defer consumer.Close()

	for {
		err = consumer.Read()
		if err != nil {
			return err
		}
	}
}

func parseConfigValues() (*kafka.ConsumerConfig, error) {

	var config confg

	_, err := conf.Parse("", &config)
	if err != nil {
		return nil, err
	}

	return &kafka.ConsumerConfig{
		BootstrapServers: config.Kafka.BootstrapServers,
		AutoOffsetReset:  config.Kafka.AutoOffsetReset,
		GroupId:          config.Kafka.GroupId,
		Topic:            config.Kafka.Topic,
	}, nil
}
