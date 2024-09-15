package main

import (
	"file-modification-tracker/daemon"
	"file-modification-tracker/http"
	"file-modification-tracker/config"
	"file-modification-tracker/ui"
	"fmt"
	"os"
)

func main() {
	config.LoadConfig()

	err := daemon.Run()
	if err != nil {
		fmt.Println("Error starting daemon:", err)
		os.Exit(1)
	}

	go http.Run()

	go ui.Run()

	select {}
}