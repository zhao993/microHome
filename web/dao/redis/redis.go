package redis

import (
	"github.com/gomodule/redigo/redis"
	"microHome/web/setting"
	"time"
)

var RedisPool *redis.Pool

func Init(cfg *setting.RedisConfig) {
	redisPool := redis.Pool{
		MaxIdle:         cfg.MaxIdle,
		MaxActive:       cfg.MaxActive,
		MaxConnLifetime: cfg.MaxConnLifetime * time.Second,
		IdleTimeout:     cfg.IdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.RedisHost+":"+cfg.RedisPort)
		},
	}
	RedisPool = &redisPool
}

func Close() {
	RedisPool.Close()
}
