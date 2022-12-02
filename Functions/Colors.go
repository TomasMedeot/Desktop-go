package functions

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Color(window fyne.Window) {
	text := canvas.NewText("", color.Black)
	inp := widget.NewEntry()
	green := widget.NewCheck("Green", func(value bool) {
	})
	red := widget.NewCheck("Red", func(value bool) {
	})
	blue := widget.NewCheck("Blue", func(value bool) {
	})
	button := widget.NewButton("Colorear", func() {
		text.Text = inp.Text
	})
	content := container.NewVBox(inp, green, red, blue, button, text)
	window.SetContent(content)
}
