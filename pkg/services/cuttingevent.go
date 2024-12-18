package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/utilis"
)

// SubscribeAndConsumeCuttingEvents interact notification of cutting result notify
func (n *NotificationService) SubScribeAnsConsumeCuttingEvents() error {
	for {
		// cuttingConsumer is similar to paymentConsumer
		msg, err := n.cuttingResultconsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer error", err)
			continue
		} else {
			fmt.Println("Message Received", msg)
		}

		var cuttingEvent models.CuttingResultEvent
		if err := json.Unmarshal(msg.Value, &cuttingEvent); err != nil {
			fmt.Printf("Failed to unmarshal cutting event message: %v\n", err)
			continue
		}
		fmt.Printf("Raw Message as String: %s\n", string(msg.Value))
		fmt.Printf("Received Cutting Event: %v\n", cuttingEvent)

		// Construct a PaymentEvent and add required details
		paymentEvent := models.PaymentEvent{
			PaymentID: string(cuttingEvent.CuttingResultID),
			OrderID:   cuttingEvent.ItemID,
			Email:     emailStore[fmt.Sprintf("%d", cuttingEvent.CuttingResultID)],
			Amount:    100.0,
			Date:      time.Now().Format("2006-01-02"),
		}

		// cuttingEvent.Components
		components := cuttingEvent.Components
		// Subject added for the notification
		subject := "Cutting Result Notification"
		body := fmt.Sprintf("Cutting Result for Order ID %d and Cutting Result ID %d", cuttingEvent.ItemID, cuttingEvent.CuttingResultID)

		// Call SendNotificationToEmail with the correct arguments
		err = utilis.SendNotificationToEmail(
			paymentEvent,                        // Pass the payment event
			subject,                             // Subject
			body,                                // Body
			cuttingEvent.CuttingResultID,        // Cutting result ID
			"Detailed cutting information here", // Cutting details
			components,                          // Pass components array here
		)

		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
		} else {
			fmt.Println("Notification sent successfully!")
		}
	}
}

// package services

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
// 	"github.com/ratheeshkumar25/opti_cut_notification/pkg/utilis"
// )

// // SubscribeAndConsumeCuttingEvents interact notification of cutting result notify
// func (n *NotificationService) SubScribeAnsConsumeCuttingEvents() error {
// 	for {
// 		// cuttingConsumer is similar to paymentConsumer
// 		msg, err := n.cuttingResultconsumer.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Println("Consumer error", err)
// 			continue
// 		} else {
// 			fmt.Println("Message Received", msg)
// 		}

// 		var cuttingEvent models.CuttingResultEvent
// 		if err := json.Unmarshal(msg.Value, &cuttingEvent); err != nil {
// 			fmt.Printf("Failed to unmarshal cutting event message: %v\n", err)
// 			continue
// 		}
// 		fmt.Printf("Raw Message as String: %s\n", string(msg.Value))
// 		fmt.Printf("Received Cutting Event: %v\n", cuttingEvent)

// 		// construct a PaymentEvent and add required details
// 		paymentEvent := models.PaymentEvent{
// 			PaymentID: string(cuttingEvent.CuttingResultID),
// 			OrderID:   cuttingEvent.ItemID,
// 			Email:     emailStore[fmt.Sprintf("%d", cuttingEvent.CuttingResultID)],
// 			Amount:    100.0,
// 			Date:      time.Now().Format("2006-01-02"),
// 		}

// 		// cuttingEvent.Components`
// 		components := cuttingEvent.Components
// 		//Subject added for the notification
// 		subject := "Cutting Result Notification"
// 		body := fmt.Sprintf("Cutting Result for Order ID %d and Cutting Result ID %d", cuttingEvent.ItemID, cuttingEvent.CuttingResultID)

// 		//call SendCuttingResultEmail with the correct arguments
// 		err = utilis.SendNotificationToEmail(
// 			paymentEvent,                        // Pass the payment event
// 			cuttingEvent.CuttingResultID,        // Cutting result ID
// 			cuttingEvent.ItemID,                 // Order ID or Item ID
// 			"Detailed cutting information here", // Cutting details, you might need to customize
// 			time.Now().Format("2006-01-02"),     // Date, formatted as string
// 			subject,                             // Subject
// 			body,                                // Body
// 			components,                          // Pass components array here
// 		)
// 		if err != nil {
// 			fmt.Printf("Error sending notification: %v\n", err)
// 		} else {
// 			fmt.Println("Notification sent successfully!")
// 		}
// 	}
// }
