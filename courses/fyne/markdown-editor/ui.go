package main

import (
	fy "fyne.io/fyne/v2"
	fyWidget "fyne.io/fyne/v2/widget"
)

type AppUI struct {
	EditWidget    *fyWidget.Entry
	PreviewWidget *fyWidget.RichText
	CurrentFile   *fy.URI
	SaveMenuItem  *fy.MenuItem
}

var ui AppUI

func (ui *AppUI) setup() {
	ui.PreviewWidget = fyWidget.NewRichTextFromMarkdown("")

	ui.EditWidget = fyWidget.NewMultiLineEntry()
	ui.EditWidget.OnChanged = ui.PreviewWidget.ParseMarkdown
}
