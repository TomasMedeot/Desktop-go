package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	App := app.New()

	Window := App.NewWindow("First Desktop")

	text := widget.NewLabel("")
	inp := widget.NewEntry()
	inp.SetPlaceHolder("ingrese un texto")
	Button := widget.NewButton("Enviar", func() {
		text.SetText(inp.Text)
	})

	content := container.NewVBox(inp, Button, text)

	Window.SetContent(content)
	Window.Resize(fyne.Size{Width: 1000.0, Height: 1000.0})

	Window.ShowAndRun()
}
