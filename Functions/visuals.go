package functions

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Test(window fyne.Window) {
	text := widget.NewLabel("")
	inp1 := widget.NewEntry()
	inp2 := widget.NewEntry()
	selec := widget.NewSelect([]string{"sum", "res"}, func(s string) {
	})
	button := widget.NewButton("Enviar", func() {
		switch selec.SelectedIndex() {
		case 0:
			text.SetText(strconv.Itoa(Calc(inp1.Text, inp2.Text, "sum")))
		case 1:
			text.SetText(strconv.Itoa(Calc(inp1.Text, inp2.Text, "res")))
		}
	})
	// button2 := widget.NewButton("Menu", func() {
	// })
	content := container.NewVBox(inp1, inp2, selec, button, text)
	window.SetContent(content)
}
