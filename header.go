package sawari

import (
	"time"

	"github.com/jung-kurt/gofpdf"
)

const (
	title  = "Fuel/Transport Expense Receipt Form"
	imageX = 10
	imageY = 60

	titleX = 60
	titleY = 30

	// HeaderBoxHeight standard height of a cell
	HeaderBoxHeight = 7
	// HeaderBoxWidth standard width of a cell
	HeaderBoxWidth = 80
)

// MakeHeader generates header
func MakeHeader(pdf *gofpdf.Fpdf, creds Creds) {

	pdf.ImageOptions("assets/10p.png", imageX, imageY, 40, 40, true, gofpdf.ImageOptions{}, 0, "")
	pdf.SetXY(titleX, titleY)
	pdf.Write(1, title)

	strPoint := float64(42)

	pdf.SetFont("Arial", "", 12)
	attrs := []string{"Name", "ID", "Designation", "Monthly Celling"}
	attrValues := []string{creds.Name, creds.EmployeeCode, creds.Designation, creds.MonthlyCeiling}
	for j, attr := range attrs {
		pdf.SetXY(25, strPoint+float64((j+1)*HeaderBoxHeight))
		pdf.Write(1, attr)
	}

	for i := 1; i <= 4; i++ {
		pdf.Rect(60, 40+float64(i*HeaderBoxHeight), float64(HeaderBoxWidth), float64(HeaderBoxHeight), "D")
		pdf.SetXY(60, 43+float64(i*HeaderBoxHeight))
		pdf.Write(1, attrValues[i-1])
	}

	makeDateBox(pdf)

	pdf.SetXY(25, 80)
	pdf.Write(1, "Note:")
	pdf.SetXY(60, 80)
	pdf.Write(1, "Don't forget to attach receipts!")

}

func makeDateBox(pdf *gofpdf.Fpdf) {

	dt := time.Now()

	pdf.Rect(140, 54, 40, float64(HeaderBoxHeight), "D")
	pdf.SetXY(140, 50)
	pdf.Write(1, "Date")
	pdf.SetXY(140, 57)
	pdf.Writef(1, "%d %s %d", dt.Day(), dt.Month(), dt.Year())

}
