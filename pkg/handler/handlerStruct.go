package handler

import (
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/services/interfaces"
)

type notifcationHandler struct {
	services interfaces.NotificationServiceInter
}

func NewNotificationHandler(service interfaces.NotificationServiceInter) *notifcationHandler {
	return &notifcationHandler{
		services: service,
	}
}
