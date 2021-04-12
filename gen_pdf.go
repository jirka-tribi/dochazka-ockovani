package main

import (
	"bytes"
	"embed"
	"github.com/signintech/gopdf"
	"io"
	"strconv"
	"time"
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

var err error

var weekDays = map[int]string{
	0: "NE",
	1: "PO",
	2: "ÚT",
	3: "ST",
	4: "ČT",
	5: "PÁ",
	6: "SO",
}

var thinkLine = 0.7
var boldLine = 1.6
var bold2Line = float64(2)

var centerAlign = gopdf.CellOption{
	Align:  gopdf.Center | gopdf.Middle,
	Border: 0,
	Float:  0,
}

var leftAlign = gopdf.CellOption{
	Align:  gopdf.Left | gopdf.Middle,
	Border: 0,
	Float:  0,
}

func fillAlignCell(pdf *gopdf.GoPdf, x float64, y float64, w float64, h float64, fillString string, align gopdf.CellOption) {
	pdf.SetX(x)
	pdf.SetY(y)
	err = pdf.CellWithOption(&gopdf.Rect{W: w, H: h}, fillString, align)
	if err != nil {
		return
	}
}

func fillHead(pdf *gopdf.GoPdf, year int, monthName string) {

	xHeadStart := float64(20)
	yHeadStart := float64(30)

	err = pdf.SetFont("timesBold", "", 16)
	if err != nil {
		return
	}

	fillAlignCell(pdf, xHeadStart, yHeadStart, 50, 16, "FN Brno", leftAlign)

	err = pdf.SetFont("times", "", 16)
	if err != nil {
		return
	}

	fillAlignCell(pdf, xHeadStart, yHeadStart+20, 50, 16, "Výkaz mzdových nároků za měsíc ", leftAlign)
	err = pdf.SetFont("timesBold", "", 16)
	if err != nil {
		return
	}
	fillAlignCell(pdf, 244, yHeadStart+20, 50, 16, monthName+" "+strconv.Itoa(year), leftAlign)

	err = pdf.SetFont("timesBold", "", 12)
	if err != nil {
		return
	}
	fillAlignCell(pdf, xHeadStart+360, yHeadStart+20, 50, 12, "NS:", leftAlign)
	fillAlignCell(pdf, xHeadStart+360, yHeadStart+45, 50, 12, "Osobní č:", leftAlign)

	pdf.SetLineWidth(bold2Line)
	pdf.Line(xHeadStart, yHeadStart+90, xHeadStart+340, yHeadStart+90)

	err = pdf.SetFont("times", "", 10)
	if err != nil {
		return
	}
	fillAlignCell(pdf, xHeadStart+25, yHeadStart+95, 50, 10, "příjmení", leftAlign)
	fillAlignCell(pdf, xHeadStart+125, yHeadStart+95, 50, 10, "jméno", leftAlign)
	fillAlignCell(pdf, xHeadStart+240, yHeadStart+95, 50, 10, "titul", leftAlign)
	fillAlignCell(pdf, xHeadStart+305, yHeadStart+95, 50, 10, "funkce", leftAlign)
}

func genPdf(year int, monthIndex int, monthName string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()

	err = pdf.AddTTFFontByReader("times", getTimesFont())
	if err != nil {
		return err
	}
	err = pdf.AddTTFFontByReader("timesBold", getTimesBoldFont())
	if err != nil {
		return err
	}

	fillHead(&pdf, year, monthName)

	xLeftStart := float64(20)
	yStart := float64(170)

	yToFill := yStart
	wColumnFirst := float64(45)
	xLeftStartAfterFirst := xLeftStart + wColumnFirst
	wColumn := float64(39)
	xRightStart := xLeftStartAfterFirst + 13*wColumn
	hHeadColumn := float64(60)
	alignHeadHelper := float64(12)
	hColumn := 15.5

	err = pdf.SetFont("times", "", 12)
	if err != nil {
		return err
	}
	pdf.SetLineWidth(thinkLine)
	pdf.RectFromUpperLeftWithStyle(xLeftStart, yToFill-30, wColumnFirst+6*wColumn, 20, "D")
	fillAlignCell(&pdf, xLeftStart+5, yToFill-30, 200, 20, "Týdenní pracovní doba:", leftAlign)

	pdf.RectFromUpperLeftWithStyle(xLeftStartAfterFirst+7*wColumn, yToFill-30, 6*wColumn, 20, "D")
	fillAlignCell(&pdf, xLeftStartAfterFirst+7*wColumn+5, yToFill-30, 200, 20, "Norma hodin v měsíci:", leftAlign)

	err = pdf.SetFont("times", "", 10)
	if err != nil {
		return err
	}

	fillAlignCell(&pdf, xLeftStart, yToFill, wColumnFirst, hHeadColumn, "Datum", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst, yToFill, wColumn, hHeadColumn-alignHeadHelper, "Rozpis", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst, yToFill+alignHeadHelper, wColumn, hHeadColumn-alignHeadHelper, "směn", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+wColumn, yToFill, wColumn, hHeadColumn, "Příchod", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+2*wColumn, yToFill, wColumn, hHeadColumn, "Odchod", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+3*wColumn, yToFill, wColumn, hHeadColumn-22, "Zúčto", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+3*wColumn, yToFill, wColumn, hHeadColumn, "vatelné", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+3*wColumn, yToFill+22, wColumn, hHeadColumn-22, "hodiny", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+4*wColumn, yToFill, 2*wColumn, hHeadColumn/2-alignHeadHelper, "z toho přes", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+4*wColumn, yToFill+alignHeadHelper, 2*wColumn, hHeadColumn/2-alignHeadHelper, "čas", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+4*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "I", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+5*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "II", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+6*wColumn, yToFill, wColumn, 47, "Noční", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+6*wColumn, yToFill+alignHeadHelper, wColumn, 47, "práce", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+7*wColumn, yToFill, wColumn, 47, "Práce v", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+7*wColumn, yToFill+alignHeadHelper, wColumn, 47, "SO a NE", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+9*wColumn, yToFill, 2*wColumn, hHeadColumn/2-alignHeadHelper, "Výkon práce při", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+9*wColumn, yToFill+alignHeadHelper, 2*wColumn, hHeadColumn/2-alignHeadHelper, "pohotovosti", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+9*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "I", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+10*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "II", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+11*wColumn, yToFill, 2*wColumn, hHeadColumn/2-alignHeadHelper, "Pohotovost mimo", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+11*wColumn, yToFill+alignHeadHelper, 2*wColumn, hHeadColumn/2-alignHeadHelper, "pracoviště", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+11*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "I", centerAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+12*wColumn, yToFill+hHeadColumn/2, wColumn, hHeadColumn/2, "II", centerAlign)

	pdf.SetLineWidth(thinkLine)
	pdf.Line(xLeftStartAfterFirst, yToFill, xLeftStartAfterFirst, yToFill+hHeadColumn)
	pdf.Line(xLeftStartAfterFirst+wColumn, yToFill, xLeftStartAfterFirst+wColumn, yToFill+hHeadColumn)
	pdf.Line(xLeftStartAfterFirst+2*wColumn, yToFill, xLeftStartAfterFirst+2*wColumn, yToFill+hHeadColumn)
	pdf.Line(xLeftStartAfterFirst+4*wColumn, yToFill+hHeadColumn/2, xLeftStartAfterFirst+6*wColumn, yToFill+hHeadColumn/2)
	pdf.Line(xLeftStartAfterFirst+9*wColumn, yToFill+hHeadColumn/2, xRightStart, yToFill+hHeadColumn/2)

	yToFill = yToFill + hHeadColumn
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	fillAlignCell(&pdf, xLeftStart, yToFill, xLeftStartAfterFirst+3*wColumn, hColumn, "Převod z minulého měsíce", centerAlign)

	xToFillNum := xLeftStartAfterFirst + 4*wColumn
	for i := 1; i <= 9; i++ {
		fillAlignCell(&pdf, xToFillNum, yToFill, wColumn, hColumn, strconv.Itoa(i), centerAlign)
		xToFillNum = xToFillNum + wColumn

	}
	yToFill = yToFill + hColumn
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)

	err = pdf.SetFont("times", "", 9)
	if err != nil {
		return err
	}

	for i := 1; i < 33; i++ {
		stringToFill := ""
		dateToGen := time.Date(2021, time.Month(monthIndex), i, 0, 0, 0, 0, time.UTC)
		if dateToGen.Month() != time.Month(monthIndex) {
			break
		}

		if int(dateToGen.Weekday()) == 6 || int(dateToGen.Weekday()) == 0 {
			pdf.SetFillColor(220, 220, 220) //setup fill color
			pdf.RectFromUpperLeftWithStyle(xLeftStart, yToFill, xRightStart-xLeftStart, hColumn, "F")
			pdf.SetFillColor(0, 0, 0)

		}

		stringToFill = strconv.Itoa(dateToGen.Day()) + "." + strconv.Itoa(monthIndex) + ".  "
		stringToFill = stringToFill + weekDays[int(dateToGen.Weekday())]
		fillAlignCell(&pdf, xLeftStart, yToFill, wColumnFirst, hColumn, stringToFill, centerAlign)
		pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)

		yToFill = yToFill + hColumn
	}

	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	fillAlignCell(&pdf, xLeftStart, yToFill, 4*wColumn, hColumn, "CELKEM", centerAlign)
	yToFill = yToFill + hColumn

	pdf.SetLineWidth(bold2Line)
	pdf.Line(xLeftStartAfterFirst+3*wColumn, yStart, xLeftStartAfterFirst+3*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+4*wColumn, yStart, xLeftStartAfterFirst+4*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+6*wColumn, yStart, xLeftStartAfterFirst+6*wColumn, yToFill)

	pdf.SetLineWidth(thinkLine)
	pdf.Line(xLeftStartAfterFirst, yStart+hHeadColumn+hColumn, xLeftStartAfterFirst, yToFill-hColumn)
	pdf.Line(xLeftStartAfterFirst+wColumn, yStart+hHeadColumn+hColumn, xLeftStartAfterFirst+wColumn, yToFill-hColumn)
	pdf.Line(xLeftStartAfterFirst+2*wColumn, yStart+hHeadColumn+hColumn, xLeftStartAfterFirst+2*wColumn, yToFill-hColumn)
	pdf.Line(xLeftStartAfterFirst+5*wColumn, yStart+hHeadColumn/2, xLeftStartAfterFirst+5*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+7*wColumn, yStart, xLeftStartAfterFirst+7*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+8*wColumn, yStart, xLeftStartAfterFirst+8*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+9*wColumn, yStart, xLeftStartAfterFirst+9*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+10*wColumn, yStart+hHeadColumn/2, xLeftStartAfterFirst+10*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+11*wColumn, yStart, xLeftStartAfterFirst+11*wColumn, yToFill)
	pdf.Line(xLeftStartAfterFirst+12*wColumn, yStart+hHeadColumn/2, xLeftStartAfterFirst+12*wColumn, yToFill)

	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	yToFill = yToFill + 2
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)

	fillAlignCell(&pdf, xLeftStart+5, yToFill, 200, hColumn, "Převod do dalšího měsíce", leftAlign)
	pdf.SetLineWidth(bold2Line)
	pdf.Line(xLeftStartAfterFirst+3*wColumn, yToFill, xLeftStartAfterFirst+3*wColumn, yToFill+hColumn)
	pdf.Line(xLeftStartAfterFirst+4*wColumn, yToFill, xLeftStartAfterFirst+4*wColumn, yToFill+hColumn)

	yToFill = yToFill + hColumn
	pdf.SetLineWidth(thinkLine)
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	fillAlignCell(&pdf, xLeftStart+5, yToFill, 200, 2*hColumn, "Zaměstnanec", leftAlign)

	yToFill = yToFill + 2*hColumn
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	fillAlignCell(&pdf, xLeftStart+5, yToFill, 200, 2*hColumn, "Vedoucí", leftAlign)
	fillAlignCell(&pdf, xLeftStartAfterFirst+5*wColumn+5, yToFill, 3*wColumn, 2*hColumn, "Personální oddělení:", leftAlign)
	pdf.Line(xLeftStartAfterFirst+5*wColumn, yToFill, xLeftStartAfterFirst+5*wColumn, yToFill+2*hColumn)

	yToFill = yToFill + 2*hColumn
	pdf.SetLineWidth(boldLine)
	pdf.Line(xLeftStart, yStart, xRightStart, yStart)
	pdf.Line(xLeftStart, yStart, xLeftStart, yToFill)
	pdf.Line(xLeftStart, yToFill, xRightStart, yToFill)
	pdf.Line(xRightStart, yStart, xRightStart, yToFill)

	err = pdf.WritePdf(monthName + "_" + strconv.Itoa(year) + ".pdf")
	if err != nil {
		return err
	}

	return nil
}
