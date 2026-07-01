package main

import (
	fy "fyne.io/fyne/v2"
	fyApp "fyne.io/fyne/v2/app"
	fyContainer "fyne.io/fyne/v2/container"
	fyWidget "fyne.io/fyne/v2/widget"
)

type App struct {
	output *fyWidget.Label
}

type appUI struct {
	output *fyWidget.Label
	entry  *fyWidget.Entry
	button *fyWidget.Button
}

var myApp App

func firstUIExample() {
	app := fyApp.New()
	window := app.NewWindow("Hello World")
	ui := myApp.initUI()

	window.SetContent(
		fyContainer.NewVBox(
			ui.output,
			ui.entry,
			ui.button,
		),
	)

	window.Resize(fy.Size{
		Width:  300,
		Height: 300,
	})

	window.ShowAndRun()
}

func (app *App) initUI() appUI {
	output := fyWidget.NewLabel("Type and Enter to see text here")
	app.output = output

	entry := fyWidget.NewEntry()

	button := fyWidget.NewButton("Enter", onButtonClick(app, entry))
	button.Importance = fyWidget.HighImportance

	return appUI{
		output: output,
		entry:  entry,
		button: button,
	}
}

func onButtonClick(app *App, entry *fyWidget.Entry) func() {
	return func() {
		app.output.SetText(entry.Text)
	}
}
