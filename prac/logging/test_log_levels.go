package main

import (
	"log"
	"os"
)

type messageType int

// INFO, WARN, ERROR, FATAL represent the logging levels
const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func init() {
	// Initialize the loggers
	// InfoLogger
	// WarningLogger
	// ErrorLogger
	// FatalLogger

	// Lutc | Lmicroseconds | Llongfile
}

func main() {
	writeLog(INFO, "this is information message!")
	writeLog(WARNING, "this is warning message!")

}

func writeLog(msgtype messageType, msg string) {

	file, err := os.OpenFile("levels.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("log-level.log not exist!")
	}

	log.SetOutput(file)

	switch msgtype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(msg)
	}

}
