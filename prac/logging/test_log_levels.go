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

// Infologger, WarningLogger, ErrorLogger, FatalLogger definition
// Init them following ...
var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {

	file, err := os.OpenFile("levels.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("levels.log not exist!")
	}

	// The `file` is a `writer` actually
	// Show different log options
	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Lmicroseconds|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL: ", log.LUTC|log.Lmicroseconds|log.Llongfile)

}

func main() {
	InfoLogger.Println("This is information message!")

	ErrorLogger.Println("This is Error message!")
	FatalLogger.Println("This is Fatal message!")

	WarningLogger.Println("Will this line show in log file?")
}

/*
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
*/
