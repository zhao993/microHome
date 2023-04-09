package test

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func redisTest() {
	//1.连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//2.操作数据库
	reply, err := conn.Do("set", "key", "value")
	// 3.回复助手函数，确定具体数据类型
	r, e := redis.String(reply, err)
	fmt.Println(r, e)
}
