package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/conf/v2"
	"github.com/dalbarracin/kafka-golang-example/internal/kafka"
)

type confg struct {
	Kafka struct {
		BootstrapServers string `conf:"default:kafka:9091"`
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

	producerConfig, err := parseConfigValues()
	if err != nil {
		return err
	}

	producer := &kafka.KafkaProducer{}

	producer.SetConfig(producerConfig)

	err = producer.Build()
	if err != nil {
		return err
	}

	defer producer.Close()

	for {

		word := "Message published at " + time.Now().Local().String()

		err = producer.Write(word)
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}

func parseConfigValues() (*kafka.ProducerConfig, error) {

	var config confg

	_, err := conf.Parse("", &config)
	if err != nil {
		return nil, err
	}

	return &kafka.ProducerConfig{
		BootstrapServers: config.Kafka.BootstrapServers,
		Topic:            config.Kafka.Topic,
	}, nil
}
