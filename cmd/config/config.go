package config

import (
	"github.com/tkanos/gonfig"
)

type KafkaConfig interface {
	Validate() error
}

func ReadConfigFromJson(filename string, kafkaConfig KafkaConfig) {

	if err := gonfig.GetConf(filename, kafkaConfig); err != nil {
		panic(err)
	}

	if err := kafkaConfig.Validate(); err != nil {
		panic(err)
	}
}
