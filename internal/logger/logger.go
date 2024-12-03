package logger

import (
	"errors"
	"log"
	"os"
)

var LogLevel int

var infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
var debugLogger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
var errorLogger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags)

var logLevelMap = map[string]int{
	"Info":  2,
	"Debug": 1,
	"Error": 0,
}

func InfoLogger(message string) {
	if LogLevel == 2 {
		infoLogger.Println(message)
	}
}

func DebugLogger(message string) {
	if LogLevel >= 1 {
		debugLogger.Println(message)
	}
}

func ErrorLogger(message string) {
	errorLogger.Println(message)
}

func SetLogLevel(value string) error {
	intLogLevel, ok := logLevelMap[value]

	if !ok {
		return errors.New("Log Level not found. Available options: Info, Debug, Error")
	}

	LogLevel = intLogLevel

	return nil
}
