package main

import (
	"bytes"
	"embed"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"image"
	"image/png"
	"strconv"
	"time"
)

//go:embed ico.png
var icon embed.FS

func getIcon() image.Image {
	dataBytes, _ := icon.ReadFile("ico.png")
	dataReader := bytes.NewReader(dataBytes)
	imageIcon, _ := png.Decode(dataReader)
	return imageIcon
}

type Months struct {
	Id   int
	Name string
}

func monthNames() []*Months {
	return []*Months{
		{1, "Leden"},
		{2, "Únor"},
		{3, "Březen"},
		{4, "Duben"},
		{5, "Květen"},
		{6, "Červen"},
		{7, "Červenec"},
		{8, "Srpen"},
		{9, "Září"},
		{10, "Říjen"},
		{11, "Listopad"},
		{12, "Prosinec"},
	}
}

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	mainWindow := new(MyMainWindow)

	var monthName *walk.ComboBox
	var year *walk.LineEdit

	now := time.Now()
	nowMonth := int(now.Month())
	nowYear := strconv.Itoa(now.Year())

	icon, _ := walk.NewIconFromImageForDPI(getIcon(), 64)

	if err := (MainWindow{
		AssignTo: &mainWindow.MainWindow,
		Title:    "Dochazka",
		Icon:     icon,
		Size:     Size{Width: 300, Height: 300},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						Text: "E&xit",
						OnTriggered: func() {
							err := mainWindow.Close()
							if err != nil {
								return
							}
						},
					},
				},
			},
			Menu{
				Text: "Help",
				Items: []MenuItem{
					Action{
						Text: "About",
						OnTriggered: func() {
							walk.MsgBox(mainWindow, "About", "Developed by Jiri Tribula", walk.MsgBoxIconInformation)
						},
					},
				},
			},
		},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Rok:",
					},
					LineEdit{
						AssignTo:  &year,
						Text:      nowYear,
						MaxLength: 4,
					},
					Label{
						Text: "Měsíc:",
					},
					ComboBox{
						AssignTo:      &monthName,
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         monthNames(),
						CurrentIndex:  nowMonth - 1,
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					PushButton{
						Text:      "Generuj Dochazku",
						Alignment: AlignHFarVCenter,
						OnClicked: func() {
							yearInt, _ := strconv.Atoi(year.Text())
							genPdf(yearInt, monthName.CurrentIndex()+1, monthName.Text())
							walk.MsgBox(mainWindow, "Done", "Hotovo: "+monthName.Text()+".pdf", walk.MsgBoxIconInformation)
						},
					},
				},
			},
		},
	}.Create()); err != nil {
	}

	mainWindow.Run()
}
