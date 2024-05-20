package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// const layoutISO = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"

var logFile *os.File

func init() {
	now := time.Now()
	date := now.Format(dateLayout)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	err = os.MkdirAll("storage/logs/", os.ModePerm)

	if err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	path := filepath.Join(cwd, "storage/logs/", date+".log")
	newFilePath := filepath.FromSlash(path)

	logFile, err = os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(logFile)
}

func Info(info string) {
	log.Printf("INFO %s\n", info)
}

func Err(_err error) {
	log.Printf("ERROR %s\n", _err)
}
