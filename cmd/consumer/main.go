package main

import (
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

	consumerConfig, err := parseConfigValues()
	if err != nil {
		panic(err)
	}

	consumer := &kafka.KafkaConsumer{}

	consumer.SetConfig(consumerConfig)

	err = consumer.Build()
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	for {
		err = consumer.Read()
		if err != nil {
			panic(err)
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
