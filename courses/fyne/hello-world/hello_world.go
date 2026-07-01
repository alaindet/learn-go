package main

import (
	"fmt"

	fyApp "fyne.io/fyne/v2/app"
	fyWidget "fyne.io/fyne/v2/widget"
)

func helloWorldExample() {
	app := fyApp.New()
	window := app.NewWindow("Hello World")
	window.SetContent(fyWidget.NewLabel("Hello WorlD"))

	// Execute stops here and a so called "run loop" starts so that the application
	// can react to user events, like clicks and keyboard
	window.ShowAndRun()
	// window.Show()
	// app.Run()

	// This is only run when quitting the application
	// You could perform cleanup logic here
	fmt.Println("Quitting the application...")
}
