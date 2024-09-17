package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("File Modification Tracker")

	startButton := widget.NewButton("Start Service", func() {

	})

	stopButton := widget.NewButton("Stop Service", func() {

	})

	logs := widget.NewMultiLineEntry()
	logs.SetText(logs.Text)

	myWindow.SetContent(container.NewVBox(startButton, stopButton))
	myWindow.ShowAndRun()
}