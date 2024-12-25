package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for NotificationServiceInter
type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) SubscribeAndConsumePaymentEvents() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockNotificationService) SubScribeAnsConsumeCuttingEvents() error {
	args := m.Called()
	return args.Error(0)
}

// Test for PaymentHandler
func TestPaymentHandler(t *testing.T) {
	mockService := new(MockNotificationService)
	mockService.On("SubscribeAndConsumePaymentEvents").Return(nil)

	handler := NewNotificationHandler(mockService)

	err := handler.PaymentHandler()

	// Assert expectations
	assert.NoError(t, err)
	mockService.AssertCalled(t, "SubscribeAndConsumePaymentEvents")
}

// Test for CuttingResultHandler
func TestCuttingResultHandler(t *testing.T) {
	mockService := new(MockNotificationService)
	mockService.On("SubScribeAnsConsumeCuttingEvents").Return(nil)

	handler := NewNotificationHandler(mockService)

	err := handler.CuttingResultHandler()

	// Assert expectations
	assert.NoError(t, err)
	mockService.AssertCalled(t, "SubScribeAnsConsumeCuttingEvents")
}

// Test for PaymentHandler with error
func TestPaymentHandler_Error(t *testing.T) {
	mockService := new(MockNotificationService)
	mockService.On("SubscribeAndConsumePaymentEvents").Return(assert.AnError)

	handler := NewNotificationHandler(mockService)

	err := handler.PaymentHandler()

	// Assert expectations
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockService.AssertCalled(t, "SubscribeAndConsumePaymentEvents")
}

// Test for CuttingResultHandler with error
func TestCuttingResultHandler_Error(t *testing.T) {
	mockService := new(MockNotificationService)
	mockService.On("SubScribeAnsConsumeCuttingEvents").Return(assert.AnError)

	handler := NewNotificationHandler(mockService)

	err := handler.CuttingResultHandler()

	// Assert expectations
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockService.AssertCalled(t, "SubScribeAnsConsumeCuttingEvents")
}
