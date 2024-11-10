package logger

import (
	"log"
	"os"
)

// Logger is a global logger instance
var Logger *log.Logger

func init() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(logFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
