package ui

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type UI struct{}

func NewUI() *UI {
    return &UI{}
}

func (u *UI) Run() error {
    myApp := app.New()
    myWindow := myApp.NewWindow("File Modification Tracker")

    startButton := widget.NewButton("Start Service", func() {
        // Start service logic
    })

    stopButton := widget.NewButton("Stop Service", func() {
        // Stop service logic
    })

    logs := widget.NewMultiLineEntry()
    logs.SetText(logs.Text)

    myWindow.SetContent(container.NewVBox(startButton, stopButton, logs))
    myWindow.ShowAndRun()
    return nil
}
