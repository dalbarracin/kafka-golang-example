package main

import (
	"time"

	"github.com/dalbarracin/kafka-golang-example/cmd/config"
	"github.com/dalbarracin/kafka-golang-example/internal/kafka"
)

func main() {

	producer := &kafka.KafkaProducer{
		Config: &kafka.ProducerConfig{},
	}

	config.ReadConfigFromJson("config.json", producer.Config)

	err := producer.Build()

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
