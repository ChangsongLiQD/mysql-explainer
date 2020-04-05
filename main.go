package main

import (
	"flag"
	"mysql-explainer/utils/config"
	"mysql-explainer/utils/log"
)

func main() {
	initConfig()
	initLogger()

	log.GetLogger().Println("testing")
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
