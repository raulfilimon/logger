package logger

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const layoutISO = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"

func Info(info string) error {
	now := time.Now()
	dateTime := now.Format(layoutISO)
	date := now.Format(dateLayout)

	cwd, _ := os.Getwd()

	err := os.MkdirAll("storage/logs/", os.ModePerm)

	if err != nil {
		return err
	}

	path := filepath.Join(cwd, "storage/logs/", date+".log")
	newFilePath := filepath.FromSlash(path)

	file, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(dateTime + "INFO: " + info + "\n")

	if err != nil {
		return err
	}

	w.Flush()
	return nil
}

func Err(_err error) error {
	now := time.Now()
	dateTime := now.Format(layoutISO)
	date := now.Format(dateLayout)

	cwd, _ := os.Getwd()

	err := os.MkdirAll("storage/logs/", os.ModePerm)

	if err != nil {
		return err
	}

	path := filepath.Join(cwd, "storage/logs/", date+".log")
	newFilePath := filepath.FromSlash(path)

	file, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, f, l, _ := runtime.Caller(1)

	log.Printf("ERROR %s %s:%d: %v", dateTime, f, l, _err)
	return nil
}
