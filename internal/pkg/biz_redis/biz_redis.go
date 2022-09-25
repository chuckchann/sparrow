package biz_redis

import (
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
	"time"
)

var redisClient *redis.Client

func Init() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	dbIndex := viper.GetInt("redis.dbIndex")
	maxConns := viper.GetInt("redis.maxConns")
	idleTimeout := viper.GetInt("redis.idleTimeout")

	cli := redis.NewClient(&redis.Options{
		Addr:        host + ":" + port,
		Password:    password,
		DB:          dbIndex,
		MaxRetries:  5,
		IdleTimeout: time.Second * time.Duration(idleTimeout),
		PoolSize:    maxConns,
	})

	_, err := cli.Ping().Result()
	if err != nil {
		panic("can not connect to redis, err:" + err.Error())
	}

	redisClient = cli
}

func GetRedis() *redis.Client {
	return redisClient
}
