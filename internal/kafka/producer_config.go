package kafka

import (
	"fmt"
)

type ProducerConfig struct {
	BootstrapServers string
	Topic            string
}

func (c *ProducerConfig) Validate() error {

	if c.BootstrapServers == "" {
		return fmt.Errorf("error: BootstrapSevers value should not be empty")
	}

	if c.Topic == "" {
		return fmt.Errorf("error: Topic value should not be empty")
	}

	return nil
}
