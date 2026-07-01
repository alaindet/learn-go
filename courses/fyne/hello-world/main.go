package main

import (
	fyne "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := fyne.New()
	window := app.NewWindow("Hello World")
	window.SetContent(widget.NewLabel("Hello WorlD"))
	window.ShowAndRun()
}
