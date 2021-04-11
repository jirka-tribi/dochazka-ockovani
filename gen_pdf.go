package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func genPdf() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("times", "times.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("times", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(255, 255, 255) //setup fill color
	pdf.RectFromUpperLeftWithStyle(10, 10, 500, 100, "FD")
	pdf.SetFillColor(0, 0, 0)

	text := "Text"
	textw, _ := pdf.MeasureTextWidth(text)
	x := 10 + (500 / 2) - (textw / 2)

	pdf.SetX(x)

	y := 10 + (100 / 2) - (float64(14) / 2)

	pdf.SetY(y)
	err = pdf.Cell(nil, "")
	if err != nil {
		return
	}

	err = pdf.WritePdf("test.pdf")
	if err != nil {
		return
	}
}
