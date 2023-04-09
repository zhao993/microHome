package redis

import (
	"encoding/json"
	"getArea/dao/mysql"
	"getArea/model"
	"github.com/gomodule/redigo/redis"
)

func GetArea() (data []model.Area, err error) {
	var areas []model.Area
	//先从缓存中取数据,使用字节切片类型接收
	conn := RedisPool.Get()
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
	if len(areaData) == 0 {
		//从Mysql中获取数据
		mysql.DB.Find(&areas)
		//写入redis数据库,将结构体序列化后存入
		conn := RedisPool.Get()
		jsonArea, _ := json.Marshal(areas)
		conn.Do("set", "areaData", jsonArea)
	} else {
		json.Unmarshal(areaData, &areas)
	}
	data = areas
	return
}
