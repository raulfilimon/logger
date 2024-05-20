package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// const layoutISO = "2006-01-02 15:04:05"
const (
	dateLayout = "2006-01-02"
	maxSize    = 1024 * 1024 * 50 // 50MB
)

var (
	logFile *os.File
	logDate time.Time
)

func init() {
	logDate = time.Now()
	openLogFile()
}

func rotateLofFileIfNeeded() {
	info, err := logFile.Stat()

	if err != nil {
		log.Fatalf("Failed to get log file info: %v", err)
		return
	}

	now := time.Now()

	if info.Size() >= maxSize || now.Day() != logDate.Day() {
		logDate = time.Now()
		openLogFile()
	}
}

func openLogFile() {
	if logFile != nil {
		logFile.Close()
	}

	date := logDate.Format(dateLayout)
	filename := fmt.Sprintf("storage/logs/%s.log", date)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logFile = file
	log.SetOutput(logFile)
}

func Info(info string) {
	rotateLofFileIfNeeded()
	log.Printf("INFO %s\n", info)
}

func Err(_err error) {
	rotateLofFileIfNeeded()
	log.Printf("ERROR %s\n", _err)
}
