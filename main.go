package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	functions "First_Desktop/Functions"
)

func main() {
	App := app.New()
	Window := App.NewWindow("First Desktop v0.1")

	calc_button := widget.NewButton("Calculator", func() {
		functions.Calculator(Window)
	})
	color_button := widget.NewButton("Color", func() {
		functions.Color(Window)
	})

	content := container.NewVBox(calc_button, color_button)
	Window.SetContent(content)
	Window.Resize(fyne.Size{Width: 500.0, Height: 200.0})
	Window.CenterOnScreen()
	Window.ShowAndRun()

}
