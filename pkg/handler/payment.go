package handler

func (n *notifcationHandler) PaymentHandler() error {
	return n.services.SubscribeAndConsumePaymentEvents()
}

func (n *notifcationHandler) CuttingResultHandler() error {
	return n.services.SubScribeAnsConsumeCuttingEvents()
}
