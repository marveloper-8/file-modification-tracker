package logs

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	file, err := os.OpenFile("file_modification_tracker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(file, "File Modification Tracker", log.Ldate|log.Ltime|log.Lshortfile)
}