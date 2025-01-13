package config

import (
	"log"

	"github.com/IBM/sarama"
)

func NewConsumerGroup(brokers []string, groupID string) (sarama.ConsumerGroup, error) {
	// Configure Sarama logger
	sarama.Logger = log.New(log.Writer(), "[sarama]", log.LstdFlags)
	// Create a new Sarama consumer group config
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // Match the Kafka server version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return consumerGroup, nil
}
