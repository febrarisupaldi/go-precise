package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
