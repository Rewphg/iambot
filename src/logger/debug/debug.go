package debug

import (
	"io/fs"
	"log"
	"os"
)

var logger *log.Logger

func InitDebugLogger() error {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return err
	}
	logger = log.New(f, "DEBUG LEVEL:", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func GetDebugLogger() *log.Logger {
	return logger
}
