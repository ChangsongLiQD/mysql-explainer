package main

import (
	"flag"
	"fmt"
	l "log"
	"mysql-explainer/explainer"
	"mysql-explainer/utils/config"
	"mysql-explainer/utils/log"
	"mysql-explainer/utils/rule"
	"mysql-explainer/utils/sql"
)

var (
	sqlFile *string
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
	sql.Init(*sqlFile)
	rules, err := config.GetRuleSetting()
	if err != nil {
		logger.Fatal(err)
	}

	r := rule.GetRule(rules)
	opt := explainer.NewOptimize()
	opt.SetRule(r)

	for query, err := sql.ReadLine(); err != nil;{
		exp := explainer.GetExplainResult(query)
		opt.SetExplain(exp)

		rec := opt.CheckRule()
		if rec.HasRecommend(){
			fmt.Println(rec.GetAllRecommend())
		}
	}
}
