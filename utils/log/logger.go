package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func InitLogger(logPath string) error{
	logPath += "/" + getFileName()
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	logger = log.New(file, "", log.LstdFlags)
	return nil
}

func getFileName() string {
	now := time.Now()
	year, month, day := now.Date()

	return fmt.Sprintf(
		"%d-%d-%d_%d:%d:%d:%d.log",
		year,
		month,
		day,
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
	)
}

func GetLogger() *log.Logger{
	return logger
}
