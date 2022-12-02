package functions

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Color(window fyne.Window) {
	color := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	text := canvas.NewText("", color)
	inp := widget.NewEntry()
	green := widget.NewCheck("Green", func(value bool) {
		if value {
			color.G = 255
		} else {
			color.G = 0
		}
	})
	red := widget.NewCheck("Red", func(value bool) {
		if value {
			color.R = 255
		} else {
			color.R = 0
		}
	})
	blue := widget.NewCheck("Blue", func(value bool) {
		if value {
			color.B = 255
		} else {
			color.B = 0
		}
	})
	button := widget.NewButton("Colorear", func() {
		text.Text = inp.Text
		text.Color = color
		text.Refresh()
	})
	content := container.NewVBox(inp, green, red, blue, button, text)
	window.SetContent(content)
}
