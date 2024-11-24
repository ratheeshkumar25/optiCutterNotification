package services

import (
	interRepo "github.com/ratheeshkumar25/opti_cut_notification/pkg/repo/interface"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/services/interfaces"
	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	Repo                  interRepo.NotificationInter
	paymentConsumer       *kafka.Reader
	cuttingResultconsumer *kafka.Reader
}

func NewNotificationService(repo interRepo.NotificationInter, paymentCon, cuttinResCon *kafka.Reader) interfaces.NotificationServiceInter {
	return &NotificationService{
		Repo:                  repo,
		paymentConsumer:       paymentCon,
		cuttingResultconsumer: cuttinResCon,
	}
}
