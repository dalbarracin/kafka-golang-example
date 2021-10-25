package main

import (
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

	producerConfig, err := parseConfigValues()
	if err != nil {
		panic(err)
	}

	producer := &kafka.KafkaProducer{}

	producer.SetConfig(producerConfig)

	err = producer.Build()
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	for {

		word := "Message published at " + time.Now().Local().String()

		err = producer.Write(word)
		if err != nil {
			panic(err)
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
