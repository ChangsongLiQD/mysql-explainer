package main

import (
	"flag"
	"io"
	l "log"
	"mysql-explainer/explainer"
	"mysql-explainer/utils/config"
	"mysql-explainer/utils/log"
	"mysql-explainer/utils/report"
	"mysql-explainer/utils/rule"
	"mysql-explainer/utils/sql"
)

var (
	sqlFile *string
	reportFile *string
	logger * l.Logger
)

func main() {
	initConfig()
	initLogger()
	initDb()
	checkSqlFile()
}

func initConfig() {
	configFile := flag.String("conf", "config/config.yaml", "Config file path")
	sqlFile = flag.String("sql", "sql/example.sql", "Checking sql file")
	reportFile = flag.String("report", "report/report.txt", "Report file")
	flag.Parse()

	err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}
}

func initLogger() {
	logFile := config.GetConfigFile()
	err := log.InitLogger(logFile)
	logger = log.GetLogger()
	if err != nil{
		panic(err)
	}
}

func initDb(){
	c := explainer.Conn{}
	err := config.GetDatabaseSetting(&c)

	if err != nil{
		panic(err)
	}
	err = explainer.InitDb(c)
	if err != nil{
		panic(err)
	}
}

func checkSqlFile(){
	report.SetFile(*reportFile)
	sql.Init(*sqlFile)
	rules, err := config.GetRuleSetting()
	if err != nil {
		logger.Fatal(err)
	}

	r := rule.GetRule(rules)
	opt := explainer.NewOptimize()
	opt.SetRule(r)

	for {
		query, err := sql.ReadLine()
		if err != nil{
			if err != io.EOF {
				logger.Fatal(err)
			}
			sql.Close()
			break
		}

		exp := explainer.GetExplainResult(query)
		opt.SetExplain(exp)

		rec := opt.CheckRule()
		if rec.HasRecommend(){
			report.Write(query, rec.GetAllRecommend())
		}
	}
}
