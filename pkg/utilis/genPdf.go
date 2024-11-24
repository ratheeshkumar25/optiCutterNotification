package utilis

import (
	"fmt"
	"log"
	"strconv"

	"github.com/phpdave11/gofpdf"
)

func GeneratePaymentInvoicePDF(paymentID string, orderID uint, amount float64, date string, filePath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetTextColor(128, 0, 128)
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Opti_Cutter_Solutions")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "HSR Layout, Bangalore, 690514")
	pdf.Ln(6)
	pdf.Cell(0, 10, "www.opticutter.com")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "INVOICE")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	pdf.CellFormat(50, 10, "Field", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Details", "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Payment ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, paymentID, "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Order ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, strconv.Itoa(int(orderID)), "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Amount:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("$%.2f", amount), "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Date:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, date, "1", 1, "L", false, 0, "")

	pdf.SetTextColor(0, 128, 0) // Green
	pdf.SetFont("", "", 10)

	pdf.CellFormat(90, 10, "Thank you for choosing. Welcome back again!", "", 0, "C", false, 0, "")
	pdf.Ln(12)
	pdf.SetFont("", "I", 8)
	pdf.CellFormat(90, 10, "This is a system-generated invoice.", "", 0, "C", false, 0, "")

	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		log.Printf("Failed to generate PDF: %v", err)
		return err
	}

	log.Println("PDF generated successfully:", filePath)
	return nil
}
