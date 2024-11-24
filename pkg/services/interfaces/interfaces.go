package interfaces

type NotificationServiceInter interface {
	SubscribeAndConsumePaymentEvents() error
	SubScribeAnsConsumeCuttingEvents() error
}
