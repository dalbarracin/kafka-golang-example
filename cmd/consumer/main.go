package main

import (
	"github.com/dalbarracin/kafka-golang-example/cmd/config"
	"github.com/dalbarracin/kafka-golang-example/internal/kafka"
)

func main() {

	consumer := &kafka.KafkaConsumer{
		Config: &kafka.ConsumerConfig{},
	}

	config.ReadConfigFromJson("config.json", consumer.Config)

	err := consumer.Build()

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
