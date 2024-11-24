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
