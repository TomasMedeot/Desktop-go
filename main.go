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
	Window := App.NewWindow("First Desktop")

	Button := widget.NewButton("Calculator", func() {
		functions.Test(Window)
	})

	content := container.NewVBox(Button)
	Window.SetContent(content)
	Window.Resize(fyne.Size{Width: 500.0, Height: 200.0})
	Window.CenterOnScreen()
	Window.ShowAndRun()

}
