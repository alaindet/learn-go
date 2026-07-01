package main

import (
	"fmt"
	"io"

	fy "fyne.io/fyne/v2"
	fyDialog "fyne.io/fyne/v2/dialog"
	fyStorage "fyne.io/fyne/v2/storage"
	fyWidget "fyne.io/fyne/v2/widget"
)

type AppUI struct {
	EditWidget    *fyWidget.Entry
	PreviewWidget *fyWidget.RichText
	CurrentFile   fy.URI
	SaveMenuItem  *fy.MenuItem
}

var ui AppUI

var allowedExt fyStorage.FileFilter = fyStorage.NewExtensionFileFilter([]string{
	".md",
})

func (ui *AppUI) setupUI() {
	ui.PreviewWidget = fyWidget.NewRichTextFromMarkdown("")

	ui.EditWidget = fyWidget.NewMultiLineEntry()
	ui.EditWidget.OnChanged = ui.PreviewWidget.ParseMarkdown
}

func (ui *AppUI) setupMenuItems(window fy.Window) {
	ui.SaveMenuItem = fy.NewMenuItem("Save", ui.createSaveHandler(window))

	window.SetMainMenu(
		fy.NewMainMenu(
			fy.NewMenu("File",
				fy.NewMenuItem("Open...", ui.createOpenHandler(window)),
				ui.SaveMenuItem,
				fy.NewMenuItem("Save as...", ui.createSaveAsHandler(window)),
			),
		),
	)
}

func (ui *AppUI) createOpenHandler(window fy.Window) func() {
	return func() {
		openDialog := fyDialog.NewFileOpen(func(read fy.URIReadCloser, err error) {
			defer read.Close()

			if err != nil {
				fyDialog.ShowError(err, window)
				return
			}

			// User canceled
			if read == nil {
				return
			}

			data, err := io.ReadAll(read)
			if err != nil {
				fyDialog.ShowError(err, window)
				return
			}

			// Update editor
			ui.EditWidget.SetText(string(data))

			// Put file URI into the app title
			ui.CurrentFile = read.URI()
			window.SetTitle(
				fmt.Sprintf("%s - %s", window.Title(), ui.CurrentFile.Name()),
			)

			// Enable save menu item
			ui.SaveMenuItem.Disabled = false
		}, window)

		openDialog.Show()
	}
}

func (ui *AppUI) createSaveHandler(window fy.Window) func() {
	return func() {
		fmt.Println("onSaveFile")
	}
}

func (ui *AppUI) createSaveAsHandler(window fy.Window) func() {
	return func() {
		saveDialog := fyDialog.NewFileSave(func(write fy.URIWriteCloser, err error) {
			defer write.Close()

			if err != nil {
				fyDialog.ShowError(err, window)
				return
			}

			// User canceled
			if write == nil {
				return
			}

			// Save file
			editorContent := []byte(ui.EditWidget.Text)
			write.Write(editorContent)

			// Put file URI into the app title
			ui.CurrentFile = write.URI()
			window.SetTitle(
				fmt.Sprintf("%s - %s", window.Title(), ui.CurrentFile.Name()),
			)

			// Enable save menu item
			ui.SaveMenuItem.Disabled = false
		}, window)

		saveDialog.Show()
	}
}
