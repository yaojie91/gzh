package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	WxKey WxKey
	Redis Redis
}

type WxKey struct {
	AppID     string
	AppSecret string
}

type Redis struct {
	Host      string
	Port      int
	Password  string
	MaxIdle   int
	MaxActive int
	Timeout   int
}

var Conf *Config

var configPath = "/etc/gzh/config"

func init() {
	initConfig()
}

func initConfig() {
	file, err := os.Open(configPath + "/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		panic(err)
	}
}
