package redis

import (
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func Init() {
	redisPool := redis.Pool{
		MaxIdle:         10,
		MaxActive:       50,
		MaxConnLifetime: 60,
		IdleTimeout:     240,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	RedisPool = &redisPool
}

func Close() {
	RedisPool.Close()
}
