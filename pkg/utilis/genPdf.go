package utilis

import (
	"fmt"
	"log"
	"strconv"

	"github.com/phpdave11/gofpdf"
	"github.com/ratheeshkumar25/opti_cut_notification/pkg/models"
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

// GenerateCuttingResultPDF generates a PDF for the cutting result with components
func GenerateCuttingResultPDF(cuttingResultID uint, orderID uint, cuttingDetails string, date string, components []models.ComponentPayload, filePath string) error {
	// Create a new PDF document
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

	// Title for the PDF
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "CUTTING RESULT")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	// Table Header
	pdf.CellFormat(50, 10, "Field", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Details", "1", 1, "L", false, 0, "")

	// Cutting Result Details
	pdf.CellFormat(50, 10, "Cutting Result ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, strconv.Itoa(int(cuttingResultID)), "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Order ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, strconv.Itoa(int(orderID)), "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Cutting Details:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, cuttingDetails, "1", 1, "L", false, 0, "")

	pdf.CellFormat(50, 10, "Date:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, date, "1", 1, "L", false, 0, "")

	// Components Table Header
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	//pdf.Cell(0, 10, "Components", "L")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 10, "Material ID", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Details", "1", 1, "L", false, 0, "")

	// Loop through components and add them to the PDF
	for _, component := range components {
		pdf.CellFormat(50, 10, strconv.Itoa(int(component.MaterialID)), "1", 0, "L", false, 0, "")
		details := fmt.Sprintf("Door: %s, Back: %s, Side: %s, Top: %s, Bottom: %s, Shelves: %s, Panel Count: %d",
			component.DoorPanel, component.BackSidePanel, component.SidePanel, component.TopPanel, component.BottomPanel, component.ShelvesPanel, component.PanelCount)
		pdf.CellFormat(100, 10, details, "1", 1, "L", false, 0, "")
	}

	pdf.SetTextColor(0, 128, 0) // Green
	pdf.SetFont("", "", 10)

	pdf.CellFormat(90, 10, "Thank you for choosing. Welcome back again!", "", 0, "C", false, 0, "")
	pdf.Ln(12)
	pdf.SetFont("", "I", 8)
	pdf.CellFormat(90, 10, "This is a system-generated document.", "", 0, "C", false, 0, "")

	// Save the PDF to the specified file path
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		log.Printf("Failed to generate PDF: %v", err)
		return err
	}

	log.Println("PDF generated successfully:", filePath)
	return nil
}
