package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Ip           string
	Port         int
	Username     string
	Password     string
	DbName       string
	ShowLog      bool
	Debuglog     bool
	DriverName   string
	MaxIdleConns int
	MaxOpenConns int
}

func GetConfig() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(path + "/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	conf := Config{}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
