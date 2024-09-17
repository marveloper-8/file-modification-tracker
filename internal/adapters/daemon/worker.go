package daemon

import (
	"log"
)

var commandQueue = make(chan string, 10)

func AddCommand(command string) {
	commandQueue <- command
}

func executeCommand(command string) {
	log.Printf("Executing command: %s", command)
}