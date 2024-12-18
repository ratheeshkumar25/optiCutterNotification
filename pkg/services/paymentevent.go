package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/utilis"
)

var emailStore = make(map[string]string)

// SubscribeAndConsumePaymentEvents implements interfaces.NotificationServiceInter.
func (n *NotificationService) SubscribeAndConsumePaymentEvents() error {
	log.Println("waiting for the event")
	for {
		msg, err := n.paymentConsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer error", err)
			continue
		} else {
			fmt.Println("Message received: ", msg)
		}

		var paymentEvent models.PaymentEvent
		if err := json.Unmarshal(msg.Value, &paymentEvent); err != nil {
			fmt.Printf("Failed to unmarshal payment message: %v\n", err)
			continue
		}
		fmt.Println("receiving produce", paymentEvent)
		emailStore[paymentEvent.PaymentID] = paymentEvent.Email

		err = utilis.SendNotificationToEmail(paymentEvent, "Payment Notification", fmt.Sprintf("Payment received: %.2f", paymentEvent.Amount), 0, "", nil)

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
// 	"log"

// 	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
// 	"github.com/ratheeshkumar25/opti_cut_notification/pkg/utilis"
// )

// var emailStore = make(map[string]string)

// // SubscribeAndConsumePaymentEvents implements interfaces.NotificationServiceInter.
// func (n *NotificationService) SubscribeAndConsumePaymentEvents() error {
// 	log.Println("waiting for the event")
// 	for {
// 		msg, err := n.paymentConsumer.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Println("Consumer error", err)
// 			continue
// 		} else {
// 			fmt.Println("Message received: ", msg)
// 		}

// 		var paymentEvent models.PaymentEvent
// 		if err := json.Unmarshal(msg.Value, &paymentEvent); err != nil {
// 			fmt.Printf("Failed to unmarshal payment message: %v\n", err)
// 			continue
// 		}
// 		fmt.Println("receiving produce", paymentEvent)
// 		emailStore[paymentEvent.PaymentID] = paymentEvent.Email

// 		err = utilis.SendNotificationToEmail(paymentEvent, "Payment Notification", fmt.Sprintf("Payment received: %.2f", paymentEvent.Amount))

// 		if err != nil {
// 			fmt.Printf("Error sending notification: %v\n", err)
// 		} else {
// 			fmt.Println("Notification sent successfully!")
// 		}
// 	}
// }
