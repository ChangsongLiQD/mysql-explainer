package report

import (
	"bufio"
	"fmt"
	l "log"
	"mysql-explainer/utils/log"
	"os"
)

var (
	buf *bufio.Writer
	inited = false
	logger *l.Logger
	file = "report/report.txt"
)

func SetFile(f string){
	file = f
}

func initBuf() {
	file, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil{
		logger = log.GetLogger()
		logger.Fatal(err)
	}

	buf = bufio.NewWriter(file)
}

func Write(query string, recs []string) (int, error) {
	if !inited{
		initBuf()
		inited = true
	}

	r := getReportText(query, recs)

	num, err := buf.WriteString(r)
	if err != nil{
		return 0, err
	}

	err = buf.Flush()
	if err != nil {
		return 0, err
	}

	return num, nil
}

func getReportText(query string, recs []string) string{
	report := query + "\n----------------------------------\n"
	for num, res := range recs{
		report += fmt.Sprintf("(%d) %s \n\n", num + 1, res)
	}
	return report
}
