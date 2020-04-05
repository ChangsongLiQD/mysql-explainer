package main

import (
	"flag"
	"mysql-explainer/explainer"
	"mysql-explainer/utils/config"
	"mysql-explainer/utils/log"
)

func main() {
	initConfig()
	initLogger()

	initDb()


}

func initConfig() {
	configFile := flag.String("conf", "config/config.yaml", "Config file path")
	flag.Parse()

	err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}
}

func initLogger() {
	logFile := config.GetConfigFile()
	err := log.InitLogger(logFile)
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
