package logger

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"time"
)

const layoutISO = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"

func Err(err_string string) error {
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
	_, err = w.WriteString(dateTime + "ERROR: " + err_string + "\n")

	if err != nil {
		return err
	}

	w.Flush()
	return nil
}

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

func NewErr(_err error) error {
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

	log.Printf(dateTime + " ERROR: " + _err.Error() + "\n")
	return nil
}
