package sql

import (
	"bufio"
	"mysql-explainer/utils/log"
	"os"
	"strings"
)

var (
	logger = log.GetLogger()
	buf *bufio.Reader
	file *os.File
)

func Init(f string){
	var err error
	file, err = os.Open(f)
	if err != nil {
		logger.Fatal(err)
	}
	buf = bufio.NewReader(file)
}

func ReadLine() (string, error){
	sql, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	sql = strings.TrimSpace(sql)
	return sql, nil
}

func Close() {
	err := file.Close()
	if err != nil {
		logger.Fatal(err)
	}
}