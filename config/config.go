package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
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
	Host        string
	Port        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Conf *Config
var RedisPool *redis.Pool
var configPath = "/etc/gzh/config"

func init() {
	initConfig()
	initRedisPool()
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

func initRedisPool() {
	redisClient := &redis.Pool{
		MaxIdle:     Conf.Redis.MaxIdle,
		MaxActive:   Conf.Redis.MaxActive,
		IdleTimeout: Conf.Redis.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", Conf.Redis.Host+":"+Conf.Redis.Port)
			if err != nil {
				return nil, err
			}
			return con, err
		},
	}
	RedisPool = redisClient
}
