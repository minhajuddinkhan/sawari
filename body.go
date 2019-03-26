package sawari

import (
	"github.com/jung-kurt/gofpdf"
)

const (
	// RectStartingX  .
	RectStartingX = 25
	// RectStartingY .
	RectStartingY = 100
)

// MakeBody makes body
func MakeBody(pdf *gofpdf.Fpdf, entries []Entry, totalAmount int) {

	width := float64(30)

	pdf.Rect(RectStartingX, RectStartingY, float64(width), float64(HeaderBoxHeight), "D")
	pdf.SetXY(RectStartingX, RectStartingY+4)
	pdf.Write(1, "Serial")

	pdf.Rect(RectStartingX+float64(width), RectStartingY, float64(width), float64(HeaderBoxHeight), "D")
	pdf.SetXY(RectStartingX+float64(width), RectStartingY+4)
	pdf.Write(1, "Date")

	pdf.Rect(RectStartingX+(float64(width)*2), RectStartingY, float64(width)*2, float64(HeaderBoxHeight), "D")
	pdf.SetXY(RectStartingX+(float64(width)*2), RectStartingY+4)
	pdf.Write(1, "Description")

	pdf.Rect(RectStartingX+(float64(width)*4), RectStartingY, float64(width), float64(HeaderBoxHeight), "D")
	pdf.SetXY(RectStartingX+(float64(width)*4), RectStartingY+4)
	pdf.Write(1, "Amount")

	startX := float64(RectStartingX)
	startY := float64(RectStartingY)
	var totalAmountX, totalAmountY float64
	for j, entry := range entries {

		X := startX
		Y := startY + float64((j+1)*(HeaderBoxHeight))
		height := float64(HeaderBoxHeight)

		//Rect for Serial
		pdf.Rect(X, Y, width, height, "D")
		pdf.SetXY(X, Y+4)
		pdf.Writef(1, "%d", entry.Serial)

		//Rect for Date
		pdf.Rect(X+width, Y, width, float64(HeaderBoxHeight), "D")
		pdf.SetXY(X+width, Y+4)
		pdf.Writef(1, "%s", entry.Date)

		//Rect for Description
		pdf.Rect(X+(width*2), Y, width*2, float64(HeaderBoxHeight), "D")
		pdf.SetXY(X+(width*2), Y+4)
		pdf.Writef(1, "%s", entry.Description)

		//Rect for Amount
		pdf.Rect(X+(width*4), Y, width, float64(HeaderBoxHeight), "D")
		pdf.SetXY(X+(width*4), Y+4)
		pdf.Writef(1, "%s", entry.Amount)

		totalAmountX = X + (width * 4)
		totalAmountY = Y + 4

	}

	pdf.SetXY(totalAmountX-30, totalAmountY+7)
	pdf.Write(1, "Total Amount")
	pdf.Rect(totalAmountX, totalAmountY+3, width, float64(HeaderBoxHeight), "D")
	pdf.SetXY(totalAmountX, totalAmountY+7)
	pdf.Writef(1, "%d", totalAmount)

	pdf.SetXY(RectStartingX, 150)
	pdf.Line(RectStartingX, totalAmountY+RectStartingY, 60, totalAmountY+RectStartingY)
	pdf.SetXY(RectStartingX, totalAmountY+105)
	pdf.Write(1, "Signature")
}
