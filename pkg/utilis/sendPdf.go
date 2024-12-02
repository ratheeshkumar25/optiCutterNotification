package utilis

import (
	"fmt"

	"github.com/ratheeshkumar25/opti_cut_notification/config"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
	"gopkg.in/gomail.v2"
)

// Send email function using GoMail
func SendNotificationToEmail(event models.PaymentEvent, subject, body string) error {
	filePath := "invoice.pdf"

	// Generate the payment invoice
	err := GeneratePaymentInvoicePDF(event.PaymentID, event.OrderID, event.Amount, event.Date, filePath)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
	m.SetHeader("To", event.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(filePath)

	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendCuttingResultEmail(event models.PaymentEvent, cuttingResultID uint, orderID uint, cuttingDetails, date, subject, body string, components []models.ComponentPayload) error {
	filePath := "cutting_result.pdf"

	// Generate the cutting result PDF, now passing the components as an argument
	err := GenerateCuttingResultPDF(cuttingResultID, orderID, cuttingDetails, date, components, filePath)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
	m.SetHeader("To", event.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(filePath)

	// Set up the SMTP dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
