package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// create log file and safe logs on it
func SetupLogging() (*os.File, error) {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	return logFile, nil
}

// clear logs
func Truncate(logfile *os.File) {
	file := logfile
	err := os.Truncate(file.Name(), 0)
	if err != nil {
		fmt.Println("error clearing log file")
	} else {
		fmt.Println("Logs cleared successfully")
	}
}
