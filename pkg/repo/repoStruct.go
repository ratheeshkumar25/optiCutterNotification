package repo

import (
	inter "github.com/ratheeshkumar25/opti_cut_notification/pkg/repo/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection             *mongo.Collection
	NotificationCollection *mongo.Collection
}

func NewMongoRepository(mongo *mongo.Database) inter.NotificationInter {
	return &MongoRepository{
		Collection:             mongo.Collection("myNotification"),
		NotificationCollection: mongo.Collection("myNotification"),
	}
}
