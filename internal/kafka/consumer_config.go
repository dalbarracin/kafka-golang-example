package kafka

import (
	"fmt"
)

type ConsumerConfig struct {
	BootstrapServers string
	AutoOffsetReset  string
	GroupId          string
	Topic            string
}

func (c ConsumerConfig) Validate() error {

	if c.BootstrapServers == "" {
		return fmt.Errorf("error: BootstrapSevers value should not be empty")
	}

	if c.AutoOffsetReset == "" {
		return fmt.Errorf("error: AutoOffsetReset value should not be empty")
	}

	if c.GroupId == "" {
		return fmt.Errorf("error: GroupId value should not be empty")
	}

	if c.Topic == "" {
		return fmt.Errorf("error: Topic value should not be empty")
	}

	return nil
}
