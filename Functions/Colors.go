package functions

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Color(window fyne.Window) {
	text := widget.NewLabel("")
	radio := widget.NewRadioGroup([]string{"Green", "Blue", "Red"}, func(value string) {
		log.Println("Radio set to", value)
	})
	content := container.NewVBox(text, radio)
	window.SetContent(content)
}
