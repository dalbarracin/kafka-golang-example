package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v2"
	"github.com/dalbarracin/kafka-golang-example/internal/kafka"
	"github.com/dalbarracin/kafka-golang-example/internal/persistence/invoice"
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

	cfg, err := parseConfigValues()
	if err != nil {
		return err
	}

	c := &kafka.KafkaConsumer{}

	c.SetConfig(cfg)

	err = c.Build()
	if err != nil {
		return err
	}

	defer c.Close()

	r, err := invoice.NewRepository()
	if err != nil {
		return err
	}

	defer r.Close()

	for {

		msg, err := c.Read()
		if err != nil {
			return err
		}

		err = r.Insert(msg)
		if err != nil {
			return err
		}
	}
}

func parseConfigValues() (*kafka.ConsumerConfig, error) {

	var cfg confg

	_, err := conf.Parse("", &cfg)
	if err != nil {
		return nil, err
	}

	return &kafka.ConsumerConfig{
		BootstrapServers: cfg.Kafka.BootstrapServers,
		AutoOffsetReset:  cfg.Kafka.AutoOffsetReset,
		GroupId:          cfg.Kafka.GroupId,
		Topic:            cfg.Kafka.Topic,
	}, nil
}
