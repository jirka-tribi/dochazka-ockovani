package main

import (
	"bytes"
	"embed"
	"github.com/signintech/gopdf"
	"io"
	"strconv"
)

//go:embed times.ttf
var timesFont embed.FS

func getTimesFont() io.Reader {
	dataBytes, _ := timesFont.ReadFile("times.ttf")
	dataReader := bytes.NewReader(dataBytes)
	return dataReader
}

//go:embed timesBold.ttf
var timesBoldFont embed.FS

func getTimesBoldFont() io.Reader {
	dataBytes, _ := timesBoldFont.ReadFile("timesBold.ttf")
	dataReader := bytes.NewReader(dataBytes)
	return dataReader
}

func genPdf(year int, monthIndex int, monthName string) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFontByReader("times", getTimesFont())
	if err != nil {
		return
	}
	err = pdf.AddTTFFontByReader("timesBold", getTimesBoldFont())
	if err != nil {
		return
	}

	err = pdf.SetFont("timesBold", "", 16)
	if err != nil {
		return
	}

	pdf.SetX(20)
	pdf.SetY(30)

	err = pdf.Cell(nil, "FN Brno")
	if err != nil {
		return
	}

	err = pdf.SetFont("times", "", 16)
	if err != nil {
		return
	}

	pdf.SetX(20)
	pdf.SetY(50)

	err = pdf.Cell(nil, "Výkaz mzdových nároků za měsíc ")
	if err != nil {
		return
	}

	err = pdf.SetFont("timesBold", "", 16)
	if err != nil {
		return
	}
	pdf.SetX(244)
	err = pdf.Cell(nil, monthName+" "+strconv.Itoa(year))
	if err != nil {
		return
	}

	err = pdf.SetFont("timesBold", "", 12)
	if err != nil {
		return
	}
	pdf.SetX(420)
	pdf.SetY(54)
	err = pdf.Cell(nil, "NS:")
	if err != nil {
		return
	}

	pdf.SetX(420)
	pdf.SetY(90)
	err = pdf.Cell(nil, "Osobní č:")
	if err != nil {
		return
	}

	pdf.SetLineWidth(2)
	pdf.SetLineType("")
	pdf.Line(20, 120, 360, 120)

	err = pdf.SetFont("times", "", 10)
	if err != nil {
		return
	}

	pdf.SetX(50)
	pdf.SetY(125)
	err = pdf.Cell(nil, "prijmeni                          jmeno                                 titul                        funkce")
	if err != nil {
		return
	}

	err = pdf.WritePdf(monthName + ".pdf")
	if err != nil {
		return
	}

}

func main() {

	genPdf(2021, 4, "Duben")
}
