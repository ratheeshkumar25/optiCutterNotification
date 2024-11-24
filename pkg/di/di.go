package di

import (
	"log"
	"sync"

	"github.com/ratheeshkumar25/opti_cut_notification/config"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/db"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/handler"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/kafka"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/repo"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/services"
)

func Init() {
	cnfg := config.LoadConfig()

	mongoClient, err := db.ConnectMongoDB(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	mongoDB := mongoClient.Database(cnfg.DBName)
	repos := repo.NewMongoRepository(mongoDB)

	// Initialize Kafka consumers
	paymentConsum, err := kafka.NewKafkaConsumer(cnfg.KafkaPort, "payment_service_group", "payment_topic")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer for payment topic: %v", err)
	}
	cuttingConsum, err := kafka.NewKafkaConsumer(cnfg.KafkaPort, "cutting_service_group", "cutting_topic")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer for cutting topic: %v", err)
	}

	// Initialize the NotificationService
	srv := services.NewNotificationService(repos, paymentConsum, cuttingConsum)
	handl := handler.NewNotificationHandler(srv)

	var wg sync.WaitGroup

	// Start payment handler in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := handl.PaymentHandler()
		if err != nil {
			log.Fatalf("Error in payment consumer: %v", err)
		}
	}()

	// Start cutting result handler in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := handl.CuttingResultHandler()
		if err != nil {
			log.Fatalf("Error in cutting consumer: %v", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
