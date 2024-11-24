package interfaces

import "github.com/ratheeshkumar25/opti_cut_notification/pkg/models"

type NotificationInter interface {
	NotificationStore(notify models.Notification) error
}
