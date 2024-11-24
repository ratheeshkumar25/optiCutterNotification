package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ratheeshkumar25/opti_cut_notification/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// func ConnectMongoDB(config *config.Config) (*mongo.Client, error) {
// 	// Configure MongoDB client options
// 	clientOptions := options.Client().
// 		ApplyURI(config.DBurl).
// 		SetServerSelectionTimeout(30 * time.Second).
// 		SetTLSConfig(&tls.Config{
// 			InsecureSkipVerify: false, // Only for debugging, not recommended for production
// 		})

// 	log.Printf("Attempting to connect to MongoDB at: %s", config.DBurl)

// 	// Context with extended timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	// Retry logic for MongoDB connection
// 	var mongoclient *mongo.Client
// 	var err error

// 	for i := 0; i < 3; i++ {
// 		mongoclient, err = mongo.Connect(ctx, clientOptions)
// 		if err == nil {
// 			break
// 		}
// 		log.Printf("Retrying MongoDB connection (%d/3): %v", i+1, err)
// 		time.Sleep(5 * time.Second)
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to MongoDB after retries: %w", err)
// 	}

// 	// Ping the database
// 	err = mongoclient.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
// 	}

// 	log.Println("MongoDB connection established")
// 	return mongoclient, nil
// }

func ConnectMongoDB(config *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.DBurl)
	log.Printf("Attempting to connect to MongoDB at: %s", config.DBurl)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("MongoDB connection established")
	return mongoclient, nil
}
