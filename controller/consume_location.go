package controller

import (
	"log"
	"tarik_consumer/kafkaandTarik2/service"

	"github.com/gin-gonic/gin"
)

func ListenForLocationUpdates(c *gin.Context) {
	// Define the Kafka brokers and group ID
	brokers := []string{"localhost:9092"}
	groupID := "location-consumer-group"
	// Start consuming location updates
	go func() {
		err := service.ConsumeLocationUpdation(brokers, groupID)
		if err != nil {
			log.Printf("Error consuming location updates: %v", err)
		}
	}()
	c.JSON(200, gin.H{
		"message": "Started listening for location updates",
	})
}
