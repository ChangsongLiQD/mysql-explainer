package config

import (
	c "go.uber.org/config"
	"os"
)

var (
	config *c.YAML
)

func Load(file string) error{
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	provider, err := c.NewYAML(c.Source(f))
	if err != nil {
		return err
	}
	config = provider
	return nil
}

func GetConfigFile() string {
	return config.Get("log").String()
}
