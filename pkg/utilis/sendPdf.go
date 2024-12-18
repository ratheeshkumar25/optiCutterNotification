package utilis

import (
	"fmt"

	"github.com/ratheeshkumar25/opti_cut_notification/config"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
	"gopkg.in/gomail.v2"
)

// Send email function using GoMail

// Send email function using GoMail
func SendNotificationToEmail(event models.PaymentEvent, subject, body string, cuttingResultID uint, cuttingDetails string, components []models.ComponentPayload) error {
	filePath := "combined_invoice.pdf" // Update to correct filename

	// Generate the payment invoice
	// Call this function to generate and send the combined PDF
	err := GenerateCombinedInvoicePDF(event.PaymentID, event.OrderID, event.Amount, event.Date, cuttingResultID, cuttingDetails, components, filePath)
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return err
	}

	m := gomail.NewMessage()
	if event.Email == "" {
		return fmt.Errorf("email address is empty for event: %+v", event)
	}
	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
	m.SetHeader("To", event.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(filePath)

	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

	// Send the email
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

// func SendNotificationToEmail(event models.PaymentEvent, subject, body string, cuttingResultID uint, cuttingDetails string, components []models.ComponentPayload) error {
// 	filePath := "invoice.pdf"

// 	// Generate the payment invoice
// 	// Call this function to generate and send the combined PDF
// 	err := GenerateCombinedInvoicePDF(event.PaymentID, event.OrderID, event.Amount, event.Date, cuttingResultID, cuttingDetails, components, "combined_invoice.pdf")
// 	if err != nil {
// 		fmt.Println("Error generating PDF:", err)
// 	}

// 	m := gomail.NewMessage()
// 	if event.Email == "" {
// 		return fmt.Errorf("email address is empty for event: %+v", event)
// 	}
// 	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
// 	m.SetHeader("To", event.Email)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)
// 	m.Attach(filePath)

// 	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Send email function using GoMail
// func SendNotificationToEmail(event models.PaymentEvent, subject, body string) error {
// 	filePath := "invoice.pdf"

// 	// Generate the payment invoice
// 	// Call this function to generate and send the combined PDF
// 	err := GenerateCombinedInvoicePDF(event.PaymentID, event.orderID, event.Amount, event.Date, cuttingResultID, cuttingDetails, components, "combined_invoice.pdf")
// 	if err != nil {
// 		fmt.Println("Error generating PDF:", err)
// 	}

// 	// err := GenerateCombinedInvoicePDF(event.PaymentID, event.OrderID, event.Amount, event.Date, filePath)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to generate PDF: %v", err)
// 	// }

// 	m := gomail.NewMessage()
// 	if event.Email == "" {
// 		return fmt.Errorf("email address is empty for event: %+v", event)
// 	}
// 	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
// 	m.SetHeader("To", event.Email)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)
// 	m.Attach(filePath)

// 	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SendCuttingResultEmail(event models.PaymentEvent, cuttingResultID uint, orderID uint, cuttingDetails, date, subject, body string, components []models.ComponentPayload) error {
// 	if event.Email == "" {
// 		return fmt.Errorf("no email found for payment event: %+v", event)
// 	}
// 	filePath := "cutting_result.pdf"

// 	// Generate the cutting result PDF
// 	err := GenerateCuttingResultPDF(cuttingResultID, orderID, cuttingDetails, date, components, filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to generate PDF: %v", err)
// 	}

// 	// Validate email address
// 	if _, err := mail.ParseAddress(event.Email); err != nil {
// 		return fmt.Errorf("invalid email address: %v", err)
// 	}

// 	// Create the email
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", config.LoadConfig().AppEmail)
// 	m.SetHeader("To", event.Email)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)
// 	m.Attach(filePath)

// 	// Set up SMTP dialer
// 	config := config.LoadConfig()
// 	d := gomail.NewDialer("smtp.gmail.com", 587, config.AppEmail, config.AppPassword)

// 	// Send email
// 	if err := d.DialAndSend(m); err != nil {
// 		return fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return nil
// }

// func SendCuttingResultEmail(event models.PaymentEvent, cuttingResultID uint, orderID uint, cuttingDetails, date, subject, body string, components []models.ComponentPayload) error {
// 	filePath := "cutting_result.pdf"

// 	// Generate the cutting result PDF, now passing the components as an argument
// 	err := GenerateCuttingResultPDF(cuttingResultID, orderID, cuttingDetails, date, components, filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to generate PDF: %v", err)
// 	}

// 	// Create a new email message
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "ratheeshgopinadhkumar@gmail.com")
// 	m.SetHeader("To", event.Email)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)
// 	m.Attach(filePath)

// 	// Set up the SMTP dialer
// 	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().AppEmail, config.LoadConfig().AppPassword)

// 	// Send the email
// 	if err := d.DialAndSend(m); err != nil {
// 		return fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return nil
// }
