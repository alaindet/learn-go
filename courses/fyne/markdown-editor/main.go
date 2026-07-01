package main

import (
	fy "fyne.io/fyne/v2"
	fyApp "fyne.io/fyne/v2/app"
	fyContainer "fyne.io/fyne/v2/container"
)

func main() {
	app := fyApp.New()
	window := app.NewWindow("Markdown Editor")
	ui.setup()

	window.SetContent(
		fyContainer.NewHSplit(
			ui.EditWidget,
			ui.PreviewWidget,
		),
	)

	window.Resize(fy.Size{Width: 800, Height: 500})
	window.CenterOnScreen()
	window.ShowAndRun()
}
