package ui

import (
	"log"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

func ShowUI() {
	var mw *walk.MainWindow
	declarative.MainWindow{
		AssignTo: &mw,
		Title:   "File Modification Tracker",
		Layout:  declarative.VBox{},
		MinSize: size.ISize{Width: 800, Height: 600},
		Children: []declarative.Widget{
			declarative.PushButton{
				Text: "Start Service",
				OnClicked: func() {
					log.Println("Start button clicked")
				},
			},
			declarative.PushButton{
				Text: "Stop Service",
				OnClicked: func() {
					log.Println("Stop button clicked")
				},
			},
		},
	}.Run()
}