package utility

import (
	"fmt"
	"github.com/go-redis/redis"
)

var ClientRedis *redis.Client

const (
	REDIS_NETWORK  = "tcp"
	REDIS_HOST     = "redis_server"
	REDIS_PORT     = "6379"
	REDIS_PASSWORD = ""
	REDIS_DB       = 0
)

func GetRedis() *redis.Client {
	if ClientRedis == nil {
		options := redis.Options{
			Network:            REDIS_NETWORK,
			Addr:               fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
			Dialer:             nil,
			OnConnect:          nil,
			Password:           REDIS_PASSWORD,
			DB:                 REDIS_DB,
			MaxRetries:         0,
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        0,
			ReadTimeout:        0,
			WriteTimeout:       0,
			PoolSize:           0,
			MinIdleConns:       0,
			MaxConnAge:         0,
			PoolTimeout:        0,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
		}
		ClientRedis = redis.NewClient(&options)
		fmt.Println(ClientRedis)
		fmt.Println("连接redis")
	}
	return ClientRedis
}

//func GetRedis() redis.Conn {
//	if rds == nil {
//		c, err := redis.Dial("tcp", "redis_server:6379")
//		if err != nil {
//			fmt.Println("connect to redis error", err)
//			return nil
//		}
//		rds = c
//	}
//	rds.
//	return rds
//}
