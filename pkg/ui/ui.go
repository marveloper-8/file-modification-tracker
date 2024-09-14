package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
)

func showUI() {
	var mw *walk.MainWindow
	var logTextEdit *walk.TextEdit

	MainWindow{
		AssignTo: &mw,
		Title: "File Modification Tracker",
		MinSize: Size{400, 300},
		Layout: VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &logTextEdit,
				ReadOnly: true,
			},
			PushButton{
				Text: "Start Service",
				OnClicked: func(){
					log.Println("Starting service...")
				},
			},
			PushButton{
				Text: "Stop Service",
				OnClicked: func(){
					log.Println("Stopping service...")
				},
			},
		},
	}.Run()
}