package repo

import (
	"context"
	"time"

	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
)

// NotificationStore implements interfaces.NotificationInter.
func (m *MongoRepository) NotificationStore(notify models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.NotificationCollection.InsertOne(ctx, notify)
	if err != nil {
		return err
	}

	return nil
}
