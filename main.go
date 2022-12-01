package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	functions "First_Desktop/elements"
)

func main() {
	fmt.Println(functions.Calc(1, 2, "sum"))
	App := app.New()
	content := container.NewVBox()
	Window := App.NewWindow("First Desktop")

	text := widget.NewLabel("")
	inp1 := widget.NewEntry()
	inp1.SetPlaceHolder("ingrese un texto")

	inp2 := widget.NewEntry()
	inp2.SetPlaceHolder("ingrese un texto")

	Button := widget.NewButton("Enviar", func() {
		text.SetText(inp1.Text)
	})

	selec := widget.NewSelect([]string{"sum", "res"}, func(s string) {
		switch s {
		case "sum":
			content = container.NewVBox(inp1, widget.NewLabel("+"), inp2, Button, text)
			Window.SetContent(content)
		case "res":
			content = container.NewVBox(inp1, Button, widget.NewLabel("-"), inp2, text)
			Window.SetContent(content)
		}
	})

	content = container.NewVBox(inp1, Button, selec, text)

	Window.SetContent(content)
	Window.Resize(fyne.Size{Width: 1000.0, Height: 1000.0})

	Window.ShowAndRun()
}
