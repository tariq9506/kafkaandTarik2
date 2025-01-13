package service

import (
	"context"
	"log"
	"tarik_consumer/kafkaandTarik2/config"
	"tarik_consumer/kafkaandTarik2/constants"

	"github.com/IBM/sarama"
)

func ConsumeLocationUpdation(brokers []string, groupID string) error {
	// Create a new consumer group
	consumerGroup, err := config.NewConsumerGroup(brokers, groupID)
	if err != nil {
		return err
	}
	defer consumerGroup.Close()
	// Define a consumer
	consumer := &LocationConsumer{}
	// Start consuming messages
	for {
		err := consumerGroup.Consume(context.Background(), []string{constants.TopicName}, consumer)
		if err != nil {
			log.Printf("Error consuming messages: %v", err)
			return err
		}
	}
}

// LocationConsumer defines a Sarama consumer group handler
type LocationConsumer struct{}

// Setup is run at the beginning of a new session

func (consumer *LocationConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session
func (consumer *LocationConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes messages
func (consumer *LocationConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received: %s\n", string(message.Value))
		session.MarkMessage(message, "")
		log.Println("--->", message)
	}
	return nil
}
