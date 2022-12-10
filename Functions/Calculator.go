package functions

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func calculate(var_1 string, var_2 string, calc string) (rest int) {

	num_1, err := strconv.Atoi(var_1)
	if err != nil {
		num_1 = 0
	}

	num_2, err := strconv.Atoi(var_2)
	if err != nil {
		num_2 = 0
	}

	switch calc {
	case "mul":
		rest = num_1 * num_2
	case "sum":
		rest = num_1 + num_2
	case "res":
		rest = num_1 - num_2
	case "div":
		rest = num_1 / num_2
	}
	return rest
}

func Calculator(window fyne.Window) {
	text := widget.NewLabel("")
	inp1 := widget.NewEntry()
	inp2 := widget.NewEntry()
	selec := widget.NewSelect([]string{"sum", "res", "div", "mul"}, func(s string) {
	})
	button := widget.NewButton("Enviar", func() {
		switch selec.SelectedIndex() {
		case 0:
			text.SetText(strconv.Itoa(calculate(inp1.Text, inp2.Text, "sum")))
		case 1:
			text.SetText(strconv.Itoa(calculate(inp1.Text, inp2.Text, "res")))
		case 2:
			text.SetText(strconv.Itoa(calculate(inp1.Text, inp2.Text, "div")))
		case 3:
			text.SetText(strconv.Itoa(calculate(inp1.Text, inp2.Text, "mul")))
		}
	})
	content := container.NewVBox(inp1, inp2, selec, button, text)
	window.SetContent(content)
}
