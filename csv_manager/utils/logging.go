package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

// SetupLogger sets up logging to a file
func SetupLogger() {
	// Set the path to your log file
	logFilePath := "C:\\Users\\paras.choudhary\\practice-app\\csv_manager\\app.log"

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to set up logger: %v", err)
	}
	Logger = log.New(file, "INFO: ", log.LstdFlags)
}
