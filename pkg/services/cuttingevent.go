package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
)

// SubscribeAndConsumeCuttingEvents implements interfaces.NotificationServiceInter.
func (n *NotificationService) SubScribeAnsConsumeCuttingEvents() error {
	for {
		// cuttingConsumer is similar to paymentConsumer
		msg, err := n.cuttingResultconsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer error", err)
			continue
		}

		var cuttingEvent models.CuttingResultEvent
		if err := json.Unmarshal(msg.Value, &cuttingEvent); err != nil {
			fmt.Printf("Failed to unmarshal cutting event message: %v\n", err)
			continue
		}
		fmt.Printf("Raw Message as String: %s\n", string(msg.Value))

		fmt.Printf("Received Cutting Event: %v\n", cuttingEvent)
	}
}
