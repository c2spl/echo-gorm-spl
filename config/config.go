package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	DB_FILENAME string
}

func GetConfig() Configuration {
	c := Configuration{}
	gonfig.GetConf("config/config.json", &c)
	return c
}
